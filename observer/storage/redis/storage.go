package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strings"
)

const keyObservers = "ATLAS_OBSERVERS"
const keyBlockNumber = "ATLAS_BLOCK_NUMBER_%d"
const keyXpub = "ATLAS_XPUB_%d"

type webHookOperation func(old []string, changes []string) []string

type Storage struct {
	client *redis.Client
}

func New(client *redis.Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) GetBlockNumber(coin uint) (int64, error) {
	key := fmt.Sprintf(keyBlockNumber, coin)
	cmd := s.client.Get(key)
	if cmd.Err() == redis.Nil {
		return 0, nil
	}
	return cmd.Int64()
}

func (s *Storage) SetBlockNumber(coin uint, num int64) error {
	key := fmt.Sprintf(keyBlockNumber, coin)
	err := s.client.Set(key, num, 0).Err()
	if err != nil {
		return errors.E(err, errors.Params{"block": num, "coin": coin}).PushToSentry()
	}
	return nil
}

func (s *Storage) SaveXpubAddresses(coin uint, addresses []string, xpub string) error {
	if len(addresses) == 0 {
		return errors.E("no addresses for xpub", errors.Params{"xpub": xpub}).PushToSentry()
	}

	a := make(map[string]interface{})
	for _, address := range addresses {
		a[address] = xpub
	}

	key := fmt.Sprintf(keyXpub, coin)
	err := s.saveHashMap(key, a)
	if err != nil {
		return err
	}
	j, err := json.Marshal(addresses)
	if err != nil {
		return errors.E(err, errors.Params{"addresses": addresses, "coin": coin, "xpub": xpub}).PushToSentry()
	}
	err = s.saveHashMap(key, map[string]interface{}{xpub: j})
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAddressFromXpub(coin uint, xpub string) ([]string, error) {
	key := fmt.Sprintf(keyXpub, coin)
	hm, err := s.getHashMap(key, xpub)
	if err != nil {
		return nil, err
	}
	if len(hm) == 0 {
		return nil, errors.E("xpub not found", errors.Params{"coin": coin, "xpub": xpub}).PushToSentry()
	}
	r, ok := hm[0].(string)
	if !ok {
		return nil, errors.E("failed to cast address list from xpub", errors.Params{"coin": coin, "xpub": xpub, "hm": hm}).PushToSentry()
	}
	var list []string
	err = json.Unmarshal([]byte(r), &list)
	if err != nil {
		return nil, errors.E(err, errors.Params{"coin": coin, "xpub": xpub, "hm": hm, "r": r}).PushToSentry()
	}
	return list, nil
}

func (s *Storage) GetXpubFromAddress(coin uint, address string) (string, error) {
	key := fmt.Sprintf(keyXpub, coin)
	r, err := s.getHashMap(key, address)
	if err != nil {
		return "", err
	}
	if len(r) == 0 || r[0] == nil {
		return "", fmt.Errorf("%d xpub not found for the address %s", coin, address)
	}
	xpub, ok := r[0].(string)
	if !ok {
		return "", errors.E("invalid type for xpub", errors.Params{"coin": coin, "address": address, "xpub": xpub}).PushToSentry()
	}
	return xpub, nil
}

func (s *Storage) Lookup(coin uint, addresses ...string) (observers []observer.Subscription, err error) {
	if len(addresses) == 0 {
		return nil, errors.E("cannot look up an empty list", errors.Params{"coin": coin}).PushToSentry()
	}

	keys := make([]string, len(addresses))
	for i, address := range addresses {
		keys[i] = key(coin, address)
	}

	kx := fmt.Sprintf(keyXpub, coin)
	xpubs, err := s.getHashMap(kx, addresses...)
	if err != nil {
		return nil, err
	}
	for i := range xpubs {
		r := xpubs[i]
		if r == nil {
			continue
		}
		if xpub, ok := r.(string); ok {
			keys[i] = key(coin, xpub)
		}
	}

	results, err := s.getHashMap(keyObservers, keys...)
	if err != nil {
		return nil, err
	}

	for i := range results {
		result := results[i]
		if result == nil {
			continue
		}
		if webhooks, ok := result.(string); ok {
			observers = append(observers, observer.Subscription{
				Coin:     coin,
				Address:  addresses[i],
				Webhooks: strings.Fields(webhooks),
			})
		}
	}
	return
}

func (s *Storage) Add(subs []observer.Subscription) error {
	return s.updateWebHooks(subs, add)
}

func (s *Storage) Delete(subs []observer.Subscription) error {
	return s.updateWebHooks(subs, remove)
}

func (s *Storage) updateWebHooks(subs []observer.Subscription, operation webHookOperation) error {
	fields := make(map[string]interface{})
	del := make([]string, 0)
	keys := make([]string, 0)
	for _, sub := range subs {
		keys = append(keys, key(sub.Coin, sub.Address))
	}

	results, err := s.getHashMap(keyObservers, keys...)
	if err != nil {
		return err
	}
	for i := range results {
		result := results[i]
		key := keys[i]
		var newWebHooks []string
		if oldWebHooks, ok := result.(string); ok && len(oldWebHooks) > 0 {
			old := strings.Fields(oldWebHooks)
			newWebHooks = operation(old, subs[i].Webhooks)
		} else {
			newWebHooks = operation(nil, subs[i].Webhooks)
		}
		if len(newWebHooks) == 0 {
			del = append(del, key)
			continue
		}
		fields[key] = strings.Join(newWebHooks, "\n")
	}
	if len(del) > 0 {
		err = s.deleteHashMapKey(keyObservers, del)
		if err != nil {
			return err
		}
	}
	if len(fields) > 0 {
		return s.saveHashMap(keyObservers, fields)
	}
	return nil
}

func (s *Storage) deleteHashMapKey(db string, fields []string) error {
	err := s.client.HDel(db, fields...).Err()
	if err != nil {
		return errors.E(err, errors.Params{"db": db, "fields": fields}).PushToSentry()
	}
	return nil
}

func (s *Storage) saveHashMap(db string, field map[string]interface{}) error {
	err := s.client.HMSet(db, field).Err()
	if err != nil {
		return errors.E(err, errors.Params{"db": db, "field": field}).PushToSentry()
	}
	return nil
}

func (s *Storage) getHashMap(db string, keys ...string) ([]interface{}, error) {
	cmd := s.client.HMGet(db, keys...)
	if err := cmd.Err(); err != nil {
		return nil, errors.E(err, errors.Params{"db": db, "keys": keys}).PushToSentry()
	}
	return cmd.Val(), nil
}

func add(old []string, changes []string) []string {
	if changes == nil {
		return old
	}
	if old == nil {
		return changes
	}

	var result []string
	for _, i := range changes {
		if !contains(old, i) {
			result = append(result, i)
		}
	}
	return append(old, result...)
}

func remove(old []string, remove []string) []string {
	n := make([]string, 0)
	if old == nil {
		return n
	}

	indices := make(map[string]bool)
	for _, r := range remove {
		indices[r] = true
	}
	for _, h := range old {
		if _, ok := indices[h]; !ok {
			n = append(n, h)
		}
	}
	return n
}

func key(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, address)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
