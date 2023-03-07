package httpx

import "net/http"

func (c *HttpClient) Head(url string) (*http.Response, error) {
	head, err := c.client.Head(url)
	if err != nil {
		return nil, err
	}
	return head, nil
}
