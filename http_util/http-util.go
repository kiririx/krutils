package http_util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func Client(timeout time.Duration) *HttpClient {
	client := http.DefaultClient
	client.Timeout = timeout
	return &HttpClient{
		client: client,
	}
}

func (c *HttpClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

func (c *HttpClient) GetJSON(url string, body map[string]string) (string, error) {
	if body != nil {
		if strings.Contains(url, "?") {
			url = url + "&"
		} else {
			url = url + "?"
		}
		for k, v := range body {
			url = url + k + "=" + v + "&"
		}
	}
	url = url[:len(url)-1]
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *HttpClient) PostJSON(url string, body map[string]any) (map[string]any, error) {
	m := make(map[string]any)
	var b []byte
	if body != nil {
		var err error
		b, err = json.Marshal(body)
		if err != nil {
			return m, err
		}
	}
	resp, err := c.client.Post(url, "application/json", strings.NewReader(string(b)))
	if err != nil {
		return m, err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (c *HttpClient) Post(url string, contentType string, body string) (*http.Response, error) {
	return c.client.Post(url, contentType, strings.NewReader(body))
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

func (c *HttpClient) Head(url string) (*http.Response, error) {
	head, err := c.client.Head(url)
	if err != nil {
		return nil, err
	}
	return head, nil
}
