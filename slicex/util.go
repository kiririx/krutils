package slicex

func Remove[T any](s []T, i int) []T {
	if i >= len(s) || i < 0 {
		return s
	}
	r := append(s[:i], s[i+1:]...)
	return r
}

func ArrayToSlice[T any](arr [...]T) []T {
	slice := make([]T, len(arr))
	for _, v := range arr {
		slice = append(slice, v)
	}
	return slice
}
