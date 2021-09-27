package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string, v interface{}) error
	Set(ctx context.Context, key string, v interface{}) error

	Del(ctx context.Context, key string) error

	WithExpiration(d time.Duration) Cache
}

type Options struct {
	Expiration time.Duration `json:"expiration"`
	URL        string        `json:"url"`
	Password   string        `json:"password"`
	Timeout    time.Duration `json:"timeout"`
}
