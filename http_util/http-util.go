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
	if strings.Contains(url, "?") {
		url = url + "&"
	} else {
		url = url + "?"
	}
	for k, v := range body {
		url = url + k + "=" + v + "&"
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

func (c *HttpClient) PostJSON(url string, body map[string]any) (string, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	resp, err := c.client.Post(url, "application/json", strings.NewReader(string(b)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *HttpClient) Post(url string, contentType string, body string) (*http.Response, error) {
	return c.client.Post(url, contentType, strings.NewReader(body))
}

func (c *HttpClient) Head(url string) (*http.Response, error) {
	head, err := c.client.Head(url)
	if err != nil {
		return nil, err
	}
	return head, nil
}
