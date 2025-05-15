package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path"
)

type postOpts struct {
	body any
}
type PostOpts func(*postOpts)

func WithBody(b any) PostOpts {
	return func(o *postOpts) {
		o.body = b
	}
}

func (c *HttpClient) Post(endpoint string, opts ...PostOpts) (*http.Response, error) {
	reqURL := c.BaseURL.JoinPath(path.Clean(endpoint))

	o := postOpts{
		body: map[string]string{},
	}
	for _, fn := range opts {
		fn(&o)
	}

	b, err := json.Marshal(o.body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	for key, value := range c.DefaultHeaders {
		req.Header.Add(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	return c.Client.Do(req)
}
