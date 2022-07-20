package mapx

import (
	"github.com/kiririx/krutils/strx"
)

func ContainsKey[K comparable, V any](m map[K]V, k K) bool {
	if _, ok := m[k]; ok {
		return true
	}
	return false
}

func GetContainedKeys[V any](s string, m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		if strx.Contains(s, k) {
			keys = append(keys, k)
		}
	}
	return keys
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
