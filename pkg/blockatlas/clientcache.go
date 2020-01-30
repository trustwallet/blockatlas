package blockatlas

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/util"
	"net/url"
	"strings"
	"time"
)

var (
	memoryCache *cache.Cache
)

func init() {
	memoryCache = cache.New(5*time.Minute, 5*time.Minute)
}

func (r *Request) PostWithCache(result interface{}, path string, body interface{}, cache time.Duration) error {
	key := r.generateKey(path, nil, body)
	err := getCache(key, result)
	if err == nil {
		return nil
	}

	err = r.Post(result, path, body)
	if err != nil {
		return err
	}
	setCache(key, result, cache)
	return err
}

func (r *Request) GetWithCache(result interface{}, path string, query url.Values, cache time.Duration) error {
	key := r.generateKey(path, query, nil)
	err := getCache(key, result)
	if err == nil {
		return nil
	}

	err = r.Get(result, path, query)
	if err != nil {
		return err
	}
	setCache(key, result, cache)
	return err
}

func getCache(key string, value interface{}) error {
	c, ok := memoryCache.Get(key)
	if !ok {
		return errors.E("validator cache: invalid cache key")
	}
	r, ok := c.([]byte)
	if !ok {
		return errors.E("validator cache: failed to cast cache to bytes")
	}
	err := json.Unmarshal(r, value)
	if err != nil {
		return errors.E(err, util.ErrNotFound).PushToSentry()
	}
	return nil
}

func setCache(key string, value interface{}, duration time.Duration) {
	b, err := json.Marshal(value)
	if err != nil {
		logger.Error(errors.E(err, "client cache cannot marshal cache object").PushToSentry())
	}
	memoryCache.Set(key, b, duration)
}

func (r *Request) generateKey(path string, query url.Values, body interface{}) string {
	var queryStr = ""
	if query != nil {
		queryStr = query.Encode()
	}
	url := strings.Join([]string{r.GetBase(path), queryStr}, "?")
	var b []byte
	if body != nil {
		b, _ = json.Marshal(body)
	}
	hash := sha1.Sum(append([]byte(url), b...))
	return base64.URLEncoding.EncodeToString(hash[:])
}
