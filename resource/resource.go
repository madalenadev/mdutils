package resource

import (
	"context"
)

type Resource interface {
	Get(ctx context.Context, endpoint string, body interface{}, data interface{}) error
	Post(ctx context.Context, endpoint string, body interface{}, data interface{}) error
	Put(ctx context.Context, endpoint string, body interface{}, data interface{}) error
	Delete(ctx context.Context, endpoint string, body interface{}, data interface{}) error
}

type Options struct {
	BaseURL string
	Header  map[string]string
}
