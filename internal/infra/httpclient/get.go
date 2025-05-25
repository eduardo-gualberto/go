package httpclient

import (
	"net/http"
	"path"
)

type getOpts struct {
	params map[string]string
}
type GetOpts func(*getOpts)

func WithParam(k string, v string) GetOpts {
	return func(o *getOpts) { o.params[k] = v }
}

func (c *HttpClient) Get(endpoint string, opts ...GetOpts) (*http.Response, error) {
	reqURL := c.BaseURL.JoinPath(path.Clean(endpoint))

	o := getOpts{params: make(map[string]string)}
	for _, fn := range opts {
		fn(&o)
	}

	query := reqURL.Query()
	for key, value := range o.params {
		query.Add(key, value)
	}
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	for key, value := range c.DefaultHeaders {
		req.Header.Add(key, value)
	}

	return c.Client.Do(req)
}
