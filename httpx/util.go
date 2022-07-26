package httpx

import (
	"github.com/kiririx/krutils/common"
	"github.com/kiririx/krutils/strx"
	"strings"
)

func GetURLWithQuery[V common.BaseType](url string, query map[string]V) string {
	if query != nil {
		if strings.Contains(url, "?") {
			url = url + "&"
		} else {
			url = url + "?"
		}
		for k, v := range query {
			url = url + k + "=" + strx.ToStr(v) + "&"
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
