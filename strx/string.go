package strx

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var randStrList = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type num interface {
	int | int32 | int8 | int64 | int16 | uint | uint16 | uint8 | uint32 | uint64 | float32 | float64
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

func ToStr(v any) string {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Struct:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return NumToStr(reflect.ValueOf(v).Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NumToStr(reflect.ValueOf(v).Uint())
	case reflect.Float32, reflect.Float64:
		return NumToStr(reflect.ValueOf(v).Float())
	// case reflect.Chan:
	// case reflect.Bool:
	// case reflect.Pointer:
	// case reflect.Array:
	// case reflect.Slice:
	// case reflect.Map:
	// case reflect.Func:
	// case reflect.Interface:
	// case reflect.Complex64:
	// case reflect.Complex128:
	// case reflect.UnsafePointer:
	default:
		return fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%v", v)
}

func SubStr(s string, start int, end int) string {
	if len(s) < end {
		return s
	}
	tmp := []rune(s)
	if start < 0 {
		return string(tmp[:end])
	}
	if end < 0 {
		return string(tmp[start:])
	}
	return string(tmp[start:end])
}

func Len(s string) int {
	return len([]rune(s))
}

func Contains[T string | []string](s string, substr ...T) bool {
	val := reflect.ValueOf(substr)
	for i := 0; i < val.Len(); i++ {
		kind := val.Index(i).Kind()
		switch kind {
		case reflect.String:
			if strings.Contains(s, val.Index(i).String()) {
				return true
			}
		case reflect.Slice:
			if Contains(s, val.Index(i).Interface().([]string)...) {
				return true
			}
		}
	}
	return false
}

func StartWith[T string | []string](s string, prefix ...T) bool {
	val := reflect.ValueOf(prefix)
	for i := 0; i < val.Len(); i++ {
		kind := val.Index(i).Kind()
		switch kind {
		case reflect.String:
			if strings.HasPrefix(s, val.Index(i).String()) {
				return true
			}
		case reflect.Slice:
			if StartWith(s, val.Index(i).Interface().([]string)...) {
				return true
			}
		}
	}
	return false
}

func EndWith(s string, suffix ...string) bool {
	val := reflect.ValueOf(suffix)
	for i := 0; i < val.Len(); i++ {
		kind := val.Index(i).Kind()
		switch kind {
		case reflect.String:
			if strings.HasSuffix(s, val.Index(i).String()) {
				return true
			}
		case reflect.Slice:
			if EndWith(s, val.Index(i).Interface().([]string)...) {
				return true
			}
		}
	}
	return false
}

func Equals[T string | []string](s string, str ...T) bool {
	val := reflect.ValueOf(str)
	for i := 0; i < val.Len(); i++ {
		kind := val.Index(i).Kind()
		switch kind {
		case reflect.String:
			if val.Index(i).String() == s {
				return true
			}
		case reflect.Slice:
			if Equals(s, val.Index(i).Interface().([]string)...) {
				return true
			}
		}
	}
	return false
}

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func RandomStrN(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randStrList[rand.Intn(n)]
	}
	return string(b)
}
