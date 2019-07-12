package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	. "github.com/trustwallet/blockatlas/observer"
)

const keyBlockNumber = "ATLAS_BLOCK_NUMBER_%d"

const luaLookupHooks = `
local key = ARGV[1] .. '-' .. ARGV[2]
local list = redis.call('hget', 'ATLAS_OBSERVERS', key)
if list == false then
	return {}
end
local items = {}
for hook in string.gmatch(list, '([^%z]*)%z') do
	items[#items + 1] = hook
end
return items
`

const luaAddHook = `
local key = ARGV[1] .. '-' .. ARGV[2]
local list = redis.call('hget', 'ATLAS_OBSERVERS', key)
if list ~= false then
	for hook in string.gmatch(list, '([^%z]*)%z') do
		if hook == ARGV[3] then
			return 'OK'
		end
	end
	list = list .. ARGV[3] .. '\0'
else
	list = ARGV[3] .. '\0'
end
redis.call('hset', 'ATLAS_OBSERVERS', key, list)
return 'OK'
`

const luaRemoveHook = `
local key = ARGV[1] .. '-' .. ARGV[2]
local list = redis.call('hget', 'ATLAS_OBSERVERS', key)
local new_list = ''
for hook in string.gmatch(list, '([^%z]*)%z') do
	if hook ~= ARGV[3] then
		new_list = new_list .. hook .. '\0'
	end
end
if new_list == '' then
	redis.call('hdel', 'ATLAS_OBSERVERS', key)
else
	redis.call('hset', 'ATLAS_OBSERVERS', key, new_list)
end
return 'OK'
`

type Client struct {
	client *redis.Client
	addSha, removeSha, lookupSha string
}

func New(client *redis.Client) (*Client, error) {
	c := &Client{
		client: client,
	}

	if err := loadScript(client, &c.addSha,luaAddHook); err != nil {
		return nil, fmt.Errorf("(add script) %s", err)
	}
	if err := loadScript(client, &c.removeSha, luaRemoveHook); err != nil {
		return nil, fmt.Errorf("(remove script) %s", err)
	}
	if err := loadScript(client, &c.lookupSha, luaLookupHooks); err != nil {
		return nil, fmt.Errorf("(lookup script) %s", err)
	}

	return c, nil
}

func loadScript(client *redis.Client, sha1 *string, script string) error {
	cmd := client.ScriptLoad(script)
	if err := cmd.Err(); err != nil {
		return err
	}
	*sha1 = cmd.Val()
	return nil
}

func (s *Client) Lookup(coin uint, addresses ...string) (subs []Subscription, err error) {
	pipe := s.client.Pipeline()
	var cmds []*redis.StringSliceCmd
	for _, address := range addresses {
		cmd := redis.NewStringSliceCmd("EVALSHA", s.lookupSha, 0, coin, address)
		if err := pipe.Process(cmd); err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}
	_, err = pipe.Exec()
	if err != nil {
		return nil, err
	}
	for i, cmd := range cmds {
		var hooks []string
		if err := cmd.ScanSlice(&hooks); err != nil {
			return nil, err
		}
		for _, hook := range hooks {
			subs = append(subs, Subscription{
				Coin:    coin,
				Address: addresses[i],
				WebHook: hook,
			})
		}
	}
	return
}

func (s *Client) Add(subs []Subscription) error {
	pipe := s.client.Pipeline()
	for _, sub := range subs {
		pipe.EvalSha(s.addSha, nil, sub.Coin, sub.Address, sub.WebHook)
	}
	_, err := pipe.Exec()
	return err
}

func (s *Client) Delete(subs []Subscription) error {
	pipe := s.client.Pipeline()
	for _, sub := range subs {
		pipe.EvalSha(s.removeSha, nil, sub.Coin, sub.Address, sub.WebHook)
	}
	_, err := pipe.Exec()
	return err
}

func (s *Client) GetBlockNumber(coin uint) (int64, error) {
	key := fmt.Sprintf(keyBlockNumber, coin)
	cmd := s.client.Get(key)
	if cmd.Err() == redis.Nil {
		return 0, nil
	}
	return cmd.Int64()
}

func (s *Client) SetBlockNumber(coin uint, num int64) error {
	key := fmt.Sprintf(keyBlockNumber, coin)
	return s.client.Set(key, num, 0).Err()
}
