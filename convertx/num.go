package convertx

import (
	"github.com/kiririx/krutils/algox"
	"strconv"
)

var RandomInt = algox.RandomInt

type Integer interface {
	int | int32 | int8 | int64 | int16 | uint | uint16 | uint8 | uint32 | uint64
}

func ToInt64() {

}

func ToInt(v string) int {
	_v, _ := strconv.Atoi(v)
	return _v
}
func ToInt8[T Integer](v T) int32 {
	return ToInt8(v)
}
