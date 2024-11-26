package convertx

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func StringToInt64(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func StringToInt(v string) (int, error) {
	return strconv.Atoi(v)
}
func StringToInt8(v string) (int8, error) {
	// 先将字符串转换为 int64
	i, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0, err
	}

	// 检查是否在 int8 范围内
	if i < -128 || i > 127 {
		return 0, errors.New("value out of range for int8")
	}

	// 转换为 int8 并返回
	return int8(i), nil
}

// IntToString 整数转字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// StringToFloat64 字符串转浮点数
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Float64ToString 浮点数转字符串
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToFloat32 字符串转 float32
func StringToFloat32(v string) (float32, error) {
	f, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

// Float32ToString float32 转字符串
func Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

// StringToBool 字符串转布尔值
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// BoolToString 布尔值转字符串
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// ToLower 字符串转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper 字符串转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// StringToTime 字符串转时间
func StringToTime(s string, layout string) (time.Time, error) {
	return time.Parse(layout, s)
}

// TimeToString 时间转字符串
func TimeToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// NowTimeString 获取当前时间的字符串表示
func NowTimeString(layout string) string {
	return time.Now().Format(layout)
}

// TimestampToTime 时间戳转时间
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// TimeToTimestamp 时间转时间戳
func TimeToTimestamp(t time.Time) int64 {
	return t.Unix()
}
