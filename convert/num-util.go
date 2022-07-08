package convert

type Integer interface {
	int | int32 | int8 | int64 | int16 | uint | uint16 | uint8 | uint32 | uint64
}

func ToInt64() {

}

func ToInt[T Integer](v T) int32 {
	return int32(v)
}
func ToInt8[T Integer](v T) int32 {
	return ToInt8(v)
}
