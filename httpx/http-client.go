package httpx

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	bodyParamNotValid = "body param is not valid"
)

type httpClient struct {
	client  *http.Client
	proxy   string
	headers map[string]string
}

type doReq struct {
	url         string
	method      string
	body        io.Reader
	contentType string
}

func Client() *httpClient {
	client := http.DefaultClient
	client.Timeout = time.Second * 4
	return &httpClient{
		client: client,
	}
}

func (c *httpClient) Timeout(duration time.Duration) *httpClient {
	c.client.Timeout = duration
	return c
}

func (c *httpClient) Proxy(proxyUrl string) *httpClient {
	uri, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}
	c.client.Transport = &http.Transport{Proxy: http.ProxyURL(uri)}
	c.proxy = proxyUrl
	return c
}

func (c *httpClient) Headers(headers map[string]string) *httpClient {
	c.headers = headers
	return c
}

func (c *httpClient) do(req *doReq) (*http.Response, error) {
	request, err := http.NewRequest(req.method, req.url, req.body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", req.contentType)
	if c.headers != nil {
		for k, v := range c.headers {
			request.Header.Set(k, v)
		}
	}
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
