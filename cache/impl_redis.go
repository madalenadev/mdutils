package cache

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/tradersclub/TCUtils/json"
)

type redisImpl struct {
	cache *redis.Client
	exp   time.Duration
}

// NewRedis cria uma implementação da interface Cache para utilizar o redis
func NewRedis(opts Options) Cache {
	optRedis := &redis.Options{
		Addr:     opts.URL,
		Password: opts.Password,
	}
	if opts.Timeout > 0 {
		optRedis.DialTimeout = opts.Timeout
	}
	return &redisImpl{
		exp:   opts.Expiration,
		cache: redis.NewClient(optRedis),
	}
}

func (c *redisImpl) Get(ctx context.Context, key string, v interface{}) error {
	cmd := c.cache.Get(ctx, key)
	err := cmd.Err()

	if err != nil {
		return err
	}

	if err := json.FromJSON([]byte(cmd.Val()), v); err != nil {
		return err
	}

	return nil
}

func (c *redisImpl) Set(ctx context.Context, key string, v interface{}) error {
	cmd := c.cache.Set(ctx, key, json.ToJSON(v), c.exp)

	err := cmd.Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *redisImpl) Del(ctx context.Context, key string) error {
	cmd := c.cache.Del(ctx, key)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (c redisImpl) WithExpiration(d time.Duration) Cache {
	c.exp = d
	return &c
}
