package convertx

import (
	"strconv"
	"strings"
	"time"
)

func StringToInt64(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func StringToInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}
func StringToInt8(v string) int8 {
	// 先将字符串转换为 int64
	i, _ := strconv.ParseInt(v, 10, 8)
	// 检查是否在 int8 范围内
	if i < -128 || i > 127 {
		return 0
	}
	// 转换为 int8 并返回
	return int8(i)
}

// IntToString 整数转字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// StringToFloat64 字符串转浮点数
func StringToFloat64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

// Float64ToString 浮点数转字符串
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToFloat32 字符串转 float32
func StringToFloat32(v string) float32 {
	f, _ := strconv.ParseFloat(v, 32)
	return float32(f)
}

// Float32ToString float32 转字符串
func Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

// StringToBool 字符串转布尔值
func StringToBool(s string) bool {
	v, _ := strconv.ParseBool(s)
	return v
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
