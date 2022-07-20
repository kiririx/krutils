package httpx

import (
	"errors"
	"fmt"
	"github.com/kiririx/krutils/jsonx"
	"io/ioutil"
	"net/http"
)

func (c *httpClient) Get(url string, query map[string]string) (*http.Response, error) {
	url = GetURLWithQuery(url, query)
	return c.do(&doReq{
		url:    url,
		method: http.MethodGet,
	})
}

func (c *httpClient) GetString(url string, query map[string]string) (string, error) {
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

func (c *httpClient) GetJSON(url string, body map[string]string) (map[string]any, error) {
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
	jsonMap, err := jsonx.JSON2Map(string(b))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("json to map failed, source: {%s} \n error: {%v}", string(b), err.Error()))
	}
	return jsonMap, nil
}
