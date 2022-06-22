package krutils

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	TimeMill = iota
	TimeMicro
	TimeNano
)

type num interface {
	int | int32 | int8 | int64 | int16 | uint | uint16 | uint8 | uint32 | uint64 | float32 | float64
}

type integer interface {
	int | int32 | int8 | int64 | int16 | uint | uint16 | uint8 | uint32 | uint64
}

func NumToStr[N num](v N) string {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatInt(int64(v), 10)
	case reflect.Float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	}
	return ""
}

func TimeToStr(v time.Time, pattern string) string {
	pattern = strings.ReplaceAll(pattern, "yyyy", "2006")
	pattern = strings.ReplaceAll(pattern, "MM", "01")
	pattern = strings.ReplaceAll(pattern, "dd", "02")
	pattern = strings.ReplaceAll(pattern, "HH", "15")
	pattern = strings.ReplaceAll(pattern, "mm", "04")
	pattern = strings.ReplaceAll(pattern, "ss", "05")
	return v.Format(pattern)
}

// TimestampToStr anything
func TimestampToStr(v int64, pattern string) string {
	var t time.Time
	if v < 10000000000 {
		t = time.Unix(v, 0)
	} else if v < 100000000000000 {
		t = time.UnixMilli(v)
	} else if v < 1000000000000000000 {
		t = time.UnixMicro(v)
	} else if v >= 1000000000000000000 {
		t = time.UnixMicro(v / 1000)
	} else {
		return ""
	}
	return TimeToStr(t, pattern)
}

func toStr(v any) string {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Struct:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Chan:
	case reflect.Bool:
	case reflect.Pointer:
	case reflect.Array:
	case reflect.Slice:
	case reflect.Map:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.UnsafePointer:

	}
	return ""
}

func ToInt(v string) int {
	return 0
}

func encodeBase64(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func decodeBase64(v string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", errors.New("decoding fail")
	}
	return string(b), nil
}

func MD5(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}
