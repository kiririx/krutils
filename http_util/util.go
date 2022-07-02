package http_util

import (
	"github.com/kiririx/krutils/common"
	"github.com/kiririx/krutils/str_util"
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
			url = url + k + "=" + str_util.ToStr(v) + "&"
		}
		url = url[:len(url)-1]
	}
	return url
}
