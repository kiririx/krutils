package tools

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Value struct {
	v     any
	check bool
}

func NewValue(v any) *Value {
	return &Value{v: v}
}

func (receiver *Value) IfNil(v any) *Value {
	if any(receiver.v) == nil {
		receiver.v = v
	}
	return receiver
}

func (receiver *Value) Int64Value() int64 {
	switch value := any(receiver.v).(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		if value > math.MaxInt64 {
			return math.MaxInt64
		}
		return int64(value)
	case float32:
		if float64(value) > math.MaxInt64 {
			return math.MaxInt64
		}
		if float64(value) < math.MinInt64 {
			return math.MinInt64
		}
		return int64(value)
	case float64:
		if value > math.MaxInt64 {
			return math.MaxInt64
		}
		if value < math.MinInt64 {
			return math.MinInt64
		}
		return int64(value)
	case string:
		if strings.Contains(value, ".") {
			f, err := strconv.ParseFloat(value, 64)
			if err != nil && receiver.check {
				panic("conversion failure")
			}
			return int64(f)
		}
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil && receiver.check {
			panic("conversion failure")
		}
		return i
	default:
		return 0
	}
}

func (receiver *Value) IntValue() int {
	v := receiver.Int64Value()
	if receiver.check && (v > math.MaxInt || v < math.MinInt) {
		panic("int overflow")
	}
	return int(v)
}

func (receiver *Value) Int8Value() int8 {
	v := receiver.Int64Value()
	if receiver.check && (v > math.MaxInt8 || v < math.MinInt8) {
		panic("int8 overflow")
	}
	return int8(v)
}

func (receiver *Value) Int16Value() int16 {
	v := receiver.Int64Value()
	if receiver.check && (v > math.MaxInt16 || v < math.MinInt16) {
		panic("int16 overflow")
	}
	return int16(v)
}

func (receiver *Value) Int32Value() int32 {
	v := receiver.Int64Value()
	if receiver.check && (v > math.MaxInt32 || v < math.MinInt32) {
		panic("int32 overflow")
	}
	return int32(v)
}
func (receiver *Value) Uint64Value() uint64 {
	switch value := any(receiver.v).(type) {
	case int:
		if value < 0 {
			return 0
		}
		return uint64(value)
	case int8:
		if value < 0 {
			return 0
		}
		return uint64(value)
	case int16:
		if value < 0 {
			return 0
		}
		return uint64(value)
	case int32:
		if value < 0 {
			return 0
		}
		return uint64(value)
	case int64:
		if value < 0 {
			return 0
		}
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case float32:
		if value < 0 || value > math.MaxUint64 {
			return 0
		}
		return uint64(value)
	case float64:
		if value < 0 || value > math.MaxUint64 {
			return 0
		}
		return uint64(value)
	case string:
		u, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return 0
		}
		return u
	default:
		return 0
	}
}

func (receiver *Value) UintValue() uint {
	v := receiver.Uint64Value()
	if receiver.check && v > math.MaxUint {
		panic("uint overflow")
	}
	return uint(v)
}

func (receiver *Value) Uint8Value() uint8 {
	v := receiver.Uint64Value()
	if receiver.check && v > math.MaxUint8 {
		panic("uint8 overflow")
	}
	return uint8(v)
}

func (receiver *Value) Uint16Value() uint16 {
	v := receiver.Uint64Value()
	if receiver.check && v > math.MaxUint16 {
		panic("uint16 overflow")
	}
	return uint16(v)
}

func (receiver *Value) Uint32Value() uint32 {
	v := receiver.Uint64Value()
	if receiver.check && v > math.MaxUint32 {
		panic("uint32 overflow")
	}
	return uint32(v)
}

func (receiver *Value) Float64Value() float64 {
	switch value := any(receiver.v).(type) {
	case int:
		return float64(value)
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	case uint:
		return float64(value)
	case uint8:
		return float64(value)
	case uint16:
		return float64(value)
	case uint32:
		return float64(value)
	case uint64:
		// if value > math.MaxFloat64 {
		// 	return math.Inf(1)
		// }
		return float64(value)
	case float32:
		return float64(value)
	case float64:
		return value
	case string:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0
		}
		return f
	case bool:
		if value {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func (receiver *Value) Float32Value() float32 {
	v := receiver.Float64Value()
	if receiver.check && (v > math.MaxFloat32 || v < -math.MaxFloat32) {
		panic("float32 overflow")
	}
	return float32(v)
}

func (receiver *Value) StringValue() string {
	return fmt.Sprintf("%v", receiver.v)
}
