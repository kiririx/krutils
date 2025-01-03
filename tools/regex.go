package tools

import (
	"reflect"
	"regexp"
	"strconv"
)

type Regex struct {
}

func (receiver *Regex) sURL(v string) bool {
	return receiver.Matched(`^(http(s?)?://)(?:[^/.\s]+\.)*.*(?:/[^/\s]+)*/?$`, v)
}

func (receiver *Regex) IsPhoneNumber(v string) bool {
	return receiver.Matched(`^1[3456789]\d{8}$`, v)
}

func (receiver *Regex) IsIdCard(v string) bool {
	return receiver.Matched(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`, v)
}

func (receiver *Regex) IsEmail(v string) bool {
	return receiver.Matched(`^[a-zA-Z0-9.!#$%&'*+/=?^_`+"`"+`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`, v)
}

func (*Regex) IsPort(v any) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		value := reflect.ValueOf(v).Int()
		return value >= 0 && value <= 65535
	case reflect.String:
		value, _ := strconv.Atoi(reflect.ValueOf(v).String())
		return value >= 0 && value <= 65535
	default:
		return false
	}
}

func (*Regex) Matched(pattern string, v string) bool {
	matched, _ := regexp.MatchString(pattern, v)
	return matched
}
