package krutils

type num interface {
	int | int32 | int8 | int64 | int16 | uint | uint16 | uint8 | uint32 | uint64 | float32 | float64
}

type Convert[A num] struct {
}

func (*Convert[A]) NumToStr(n A) string {
	return "123"
}
