package regx

import (
	"reflect"
	"regexp"
	"strconv"
)

// IsURL return true if the string is a URL
func IsURL(v string) bool {
	return Matched(`^(http(s?)?://)(?:[^/.\s]+\.)*.*(?:/[^/\s]+)*/?$`, v)
}

func IsPhoneNumber(v string) bool {
	return Matched(`^1[3456789]\d{8}$`, v)
}

func IsIdCard(v string) bool {
	return Matched(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`, v)
}

func IsEmail(v string) bool {
	return Matched(`^[a-zA-Z0-9.!#$%&'*+/=?^_`+"`"+`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`, v)
}

func IsPort[T string | int](v T) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		return reflect.ValueOf(v).Int() > 0
	case reflect.String:
		p, _ := strconv.Atoi(reflect.ValueOf(v).String())
		return p > 0
	}
	return false
}

func Matched(pattern string, v string) bool {
	matched, _ := regexp.MatchString(pattern, v)
	return matched
}
