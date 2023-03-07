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

type HttpClient struct {
	client  *http.Client
	proxy   string
	headers map[string]string
	cookies []*http.Cookie
}

type doReq struct {
	url         string
	method      string
	body        io.Reader
	contentType string
}

func Client() *HttpClient {
	client := http.DefaultClient
	client.Timeout = time.Second * 4
	return &HttpClient{
		client: client,
	}
}

func (c *HttpClient) Timeout(duration time.Duration) *HttpClient {
	c.client.Timeout = duration
	return c
}

func (c *HttpClient) Proxy(proxyUrl string) *HttpClient {
	uri, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}
	c.client.Transport = &http.Transport{Proxy: http.ProxyURL(uri)}
	c.proxy = proxyUrl
	return c
}

func (c *HttpClient) Headers(headers map[string]string) *HttpClient {
	c.headers = headers
	return c
}

func (c *HttpClient) Cookies(cookies []*http.Cookie) *HttpClient {
	c.cookies = cookies
	return c
}

func (c *HttpClient) do(req *doReq) (*http.Response, error) {
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
	if c.cookies != nil {
		for _, v := range c.cookies {
			request.AddCookie(v)
		}
	}
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
