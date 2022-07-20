package slicex

func Remove[T any](s []T, i int) []T {
	if i >= len(s) || i < 0 {
		return s
	}
	r := append(s[:i], s[i+1:]...)
	return r
}
