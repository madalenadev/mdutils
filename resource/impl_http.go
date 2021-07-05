package resource

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/tradersclub/TCUtils/tcerr"
)

func NewHTTP(opts Options) Resource {
	if opts.BaseURL != "" && !strings.HasSuffix(opts.BaseURL, "/") {
		opts.BaseURL += "/"
	}

	if opts.Header == nil {
		opts.Header = map[string]string{
			"Content-Type": "application/json",
		}
	}

	return &implResource{
		baseURL: opts.BaseURL,
		header:  opts.Header,
	}
}

type implResource struct {
	baseURL string
	header  map[string]string
}

func (i *implResource) request(ctx context.Context, method, endpoint string, body interface{}, data interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		bt, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(bt)
	}

	if strings.HasPrefix(endpoint, "/") {
		endpoint = endpoint[1 : len(endpoint)-1]
	}

	url := i.baseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return err
	}

	for key, value := range i.header {
		req.Header.Set(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode < 200 || res.StatusCode > 299 {
		detail := make(map[string]interface{})
		if err := decoder.Decode(&detail); err != nil {
			return err
		}

		return tcerr.NewError(res.StatusCode, http.StatusText(res.StatusCode), detail)
	}

	if data != nil {
		err = decoder.Decode(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *implResource) Get(ctx context.Context, endpoint string, body interface{}, data interface{}) error {
	return i.request(ctx, http.MethodGet, endpoint, body, data)
}

func (i *implResource) Post(ctx context.Context, endpoint string, body interface{}, data interface{}) error {
	return i.request(ctx, http.MethodPost, endpoint, body, data)
}

func (i *implResource) Put(ctx context.Context, endpoint string, body interface{}, data interface{}) error {
	return i.request(ctx, http.MethodPut, endpoint, body, data)
}

func (i *implResource) Delete(ctx context.Context, endpoint string, body interface{}, data interface{}) error {
	return i.request(ctx, http.MethodDelete, endpoint, body, data)
}
