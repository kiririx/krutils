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
