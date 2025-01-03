package tools

import (
	"strings"
)

func NewString() *String {
	return &String{}
}

type String struct {
}

// SubStrWithRune 是一个用于从字符串中提取子字符串的函数。
// 它使用 rune 作为处理单位，可以正确处理多字节字符。
// 参数:
//
//	s - 输入的字符串。
//	start - 子字符串的起始位置，0开始, 负值表示从字符串末尾开始计数。
//	end - 子字符串的结束位置，len-1为最大值, 负值表示从字符串末尾开始计数。
//
// 返回值:
//
//	返回提取的子字符串。
//
// 功能描述:
//  1. 如果输入字符串的长度小于指定的结束位置，则直接返回原字符串。
//  2. 将字符串转换为 rune 切片以处理多字节字符。
//  3. 根据 start 和 end 参数的不同情况，提取相应的子字符串。
func (receiver *String) SubStrWithRune(s string, start int, end int) string {
	// 检查字符串长度是否小于结束位置，如果是，则直接返回原字符串
	if len(s) < end {
		return s
	}
	// 将字符串转换为 rune 切片，以便正确处理多字节字符
	tmp := []rune(s)
	// 如果起始位置小于0，表示从字符串末尾开始计数，提取到 end 位置的子字符串
	if start < 0 {
		return string(tmp[:end])
	}
	// 如果结束位置小于0，表示从字符串末尾开始计数，提取从 start 位置到字符串末尾的子字符串
	if end < 0 {
		return string(tmp[start:])
	}
	// 提取从 start 到 end 位置之间的子字符串
	return string(tmp[start:end])
}

func (receiver *String) LenWithRune(s string) int {
	return len([]rune(s))
}

func (receiver *String) Contains(s string, substr ...string) bool {
	for _, s2 := range substr {
		if strings.Contains(s, s2) {
			return true
		}
	}
	return false
}

func (receiver *String) StartWith(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func (receiver *String) EndWith(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func (receiver *String) In(s string, stringArray ...string) bool {
	for _, s2 := range stringArray {
		if s == s2 {
			return true
		}
	}
	return false
}

// DefaultIfEmpty 如果字符串 s 为空，则返回 elseStr。
// 这个方法的主要作用是提供一个安全的方式来处理空字符串，
// 它允许调用者指定一个默认值，当输入的字符串为空时返回。
// 参数:
//
//	s - 需要检查的字符串。
//	elseStr - 如果 s 为空时返回的默认字符串。
//
// 返回值:
//
//	如果 s 为空，返回 elseStr；否则，返回 s。
func (receiver *String) DefaultIfEmpty(s string, defStr string) string {
	if s == "" {
		return defStr
	}
	return s
}
