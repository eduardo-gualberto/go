package httpclient

import (
	"net/http"
	"net/url"
)

type httpOpts struct {
	headers map[string]string
	baseUrl string
}
type HttpOpts func(*httpOpts)

func WithHeader(k string, v string) HttpOpts {
	return func(o *httpOpts) {
		o.headers[k] = v
	}
}

func WithBaseUrl(u string) HttpOpts {
	return func(o *httpOpts) { o.baseUrl = u }
}

type HttpClient struct {
	BaseURL        *url.URL
	Client         *http.Client
	DefaultHeaders map[string]string
}

func NewHttpClient(opts ...HttpOpts) (*HttpClient, error) {
	o := httpOpts{
		baseUrl: "",
		headers: make(map[string]string),
	}
	for _, fn := range opts {
		fn(&o)
	}
	url, err := url.Parse(o.baseUrl)
	if err != nil {
		return nil, err
	}
	client := &HttpClient{
		BaseURL:        url,
		Client:         &http.Client{},
		DefaultHeaders: o.headers,
	}
	return client, nil
}
