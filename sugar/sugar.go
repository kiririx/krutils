package sugar

func Then[T any](cond bool, v1 T, v2 T) T {
	if cond {
		return v1
	}
	return v2
}

func ThenFunc[T any](cond bool, v1 func() T, v2 func() T) T {
	if cond {
		return v1()
	}
	return v2()
}

func ForIndex(end int, f func(i int) (bool, bool)) {
	for i := 0; i < end; i++ {
		_continue, _break := f(i)
		if _break {
			break
		}
		if _continue {
			continue
		}
	}
}

func ForSlice[T any](s []T, f func(i int, v T) (bool, bool)) {
	for i, v := range s {
		_continue, _break := f(i, v)
		if _break {
			break
		}
		if _continue {
			continue
		}
	}
}
