package cache

import (
	"context"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/tradersclub/TCUtils/logger"
)

type memcacheImpl struct {
	cache *memcache.Client
	exp   time.Duration
}

// NewMemcache cria uma implementação da interface Cache para utilizar o memcache
func NewMemcache(opts Options) Cache {
	client := memcache.New(opts.URL)
	if opts.Timeout > 0 {
		client.Timeout = opts.Timeout
	}

	if err := client.Ping(); err != nil {
		logger.Fatal(err)
	}

	logger.Info("Memcached connected")
	return &memcacheImpl{
		exp:   opts.Expiration,
		cache: client,
	}
}

func (c *memcacheImpl) Get(ctx context.Context, key string, v interface{}) error {
	item, err := c.cache.Get(key)
	if err != nil {
		return err
	}

	if err := decode(item.Value, v); err != nil {
		return err
	}

	return nil
}

func (c *memcacheImpl) Set(ctx context.Context, key string, v interface{}) error {
	value, err := encode(v)
	if err != nil {
		return err
	}

	err = c.cache.Set(&memcache.Item{Key: key, Value: value, Expiration: int32(c.exp)})
	if err != nil {
		return err
	}

	return nil
}

func (c *memcacheImpl) Del(ctx context.Context, key string) error {
	err := c.cache.Delete(key)
	return err
}

func (c memcacheImpl) WithExpiration(d time.Duration) Cache {
	c.exp = d / time.Second
	return &c
}
