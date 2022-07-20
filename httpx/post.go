package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kiririx/krutils/jsonx"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *httpClient) PostString(url string, body map[string]any) (string, error) {
	bodyParams, err := jsonx.Map2JSON(body)
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

func (c *httpClient) PostJSON(url string, body map[string]any) (map[string]any, error) {
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
	jsonMap, err := jsonx.JSON2Map(string(b))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("json to map failed, source: {%s} \n error: {%v}", string(b), err.Error()))
	}
	return jsonMap, nil
}

func (c *httpClient) Post(url string, body map[string]any) (*http.Response, error) {
	bodyParams, err := jsonx.Map2JSON(body)
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

func (c *httpClient) PostFormGetJSON(url string, data map[string]string) (map[string]any, error) {
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
