package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

func NewHttpClient() *HttpClient {
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

func (c *HttpClient) Head(url string) (*http.Response, error) {
	head, err := c.client.Head(url)
	if err != nil {
		return nil, err
	}
	return head, nil
}

func (c *HttpClient) Get(url string, query map[string]string) (*http.Response, error) {
	url = GetURLWithQuery(url, query)
	return c.do(&doReq{
		url:    url,
		method: http.MethodGet,
	})
}

func (c *HttpClient) GetString(url string, query map[string]string) (string, error) {
	url = GetURLWithQuery(url, query)
	resp, err := c.do(&doReq{
		url:         url,
		method:      http.MethodGet,
		contentType: "application/json",
	})
	if err != nil {
		return "", errors.New(fmt.Sprintf("http get failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("read response body failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	return string(b), nil
}

func (c *HttpClient) GetJSON(url string, body map[string]string) (map[string]any, error) {
	url = GetURLWithQuery(url, body)
	resp, err := c.do(&doReq{
		url:         url,
		method:      http.MethodGet,
		contentType: "application/json",
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http get failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("read response body failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	jsonMap, err := NewJSON().JSON2Map(string(b))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("json to map failed, source: {%s} \n error: {%v}", string(b), err.Error()))
	}
	return jsonMap, nil
}

func (c *HttpClient) PostString(url string, body map[string]any) (string, error) {
	bodyParams, err := NewJSON().Map2JSON(body)
	if err != nil {
		return "", errors.New(bodyParamNotValid)
	}
	resp, err := c.do(&doReq{
		url:    url,
		method: http.MethodPost,
		body:   strings.NewReader(bodyParams),
	})
	if err != nil {
		return "", errors.New(fmt.Sprintf("http post failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("read response body failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	return string(b), nil
}

func (c *HttpClient) PostJSON(url string, body map[string]any) (map[string]any, error) {
	var b []byte
	if body != nil {
		var err error
		b, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	resp, err := c.do(&doReq{
		url:         url,
		method:      http.MethodPost,
		contentType: "application/json",
		body:        strings.NewReader(string(b)),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http post failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("read response body failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	jsonMap, err := NewJSON().JSON2Map(string(b))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("json to map failed, source: {%s} \n error: {%v}", string(b), err.Error()))
	}
	return jsonMap, nil
}

func (c *HttpClient) Post(url string, body map[string]any) (*http.Response, error) {
	bodyParams, err := NewJSON().Map2JSON(body)
	if err != nil {
		return nil, errors.New(bodyParamNotValid)
	}
	resp, err := c.do(&doReq{
		url:    url,
		method: http.MethodPost,
		body:   strings.NewReader(bodyParams),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http post failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	return resp, nil
}

func (c *HttpClient) PostFormGetJSON(url string, data map[string]string) (map[string]any, error) {
	resultMap := make(map[string]any)
	m := make(map[string][]string)
	for k, v := range data {
		m[k] = []string{v}
	}
	resp, err := c.client.PostForm(url, m)
	if err != nil || resp.StatusCode != http.StatusOK {
		return resultMap, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resultMap, err
	}
	err = json.Unmarshal(b, &resultMap)
	if err != nil {
		return resultMap, err
	}
	return resultMap, nil
}

func (c *HttpClient) PostStringGetJSON(url string, raw string) (map[string]any, error) {
	var b []byte
	resp, err := c.do(&doReq{
		url:         url,
		method:      http.MethodPost,
		contentType: "application/json",
		body:        strings.NewReader(raw),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http post failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("read response body failed, url: {%s}, cause: {%v} ", url, err.Error()))
	}
	jsonMap, err := NewJSON().JSON2Map(string(b))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("json to map failed, source: {%s} \n error: {%v}", string(b), err.Error()))
	}
	return jsonMap, nil
}

func GetURLWithQuery[V any](url string, query map[string]V) string {
	if query != nil {
		if strings.Contains(url, "?") {
			url = url + "&"
		} else {
			url = url + "?"
		}
		for k, v := range query {
			url = url + k + "=" + NewValue(v).StringValue() + "&"
		}
		url = url[:len(url)-1]
	}
	return url
}

func GetQueryParam(url string, key string) string {
	if strings.Contains(url, "?") {
		url = url[strings.Index(url, "?")+1:]
		var kvs []string
		if strings.Contains(url, "&") {
			kvs = strings.Split(url, "&")
		} else {
			kvs = []string{url}
		}
		for _, kv := range kvs {
			if strings.Contains(kv, "=") {
				if strings.Split(kv, "=")[0] == key {
					return strings.Split(kv, "=")[1]
				}
			}
		}
	}
	return ""
}

func GetQueryParams(url string) map[string]string {
	params := make(map[string]string)
	if strings.Contains(url, "?") {
		url = url[strings.Index(url, "?")+1:]
		var kvs []string
		if strings.Contains(url, "&") {
			kvs = strings.Split(url, "&")
		} else {
			kvs = []string{url}
		}
		for _, kv := range kvs {
			if strings.Contains(kv, "=") {
				params[strings.Split(kv, "=")[0]] = strings.Split(kv, "=")[1]
			} else {
				params[kv] = ""
			}
		}
	}
	return params
}
