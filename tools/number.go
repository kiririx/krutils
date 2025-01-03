package tools

type Number struct {
}

func NewNumber() *Number {
	return &Number{}
}

func (n *Number) DefaultInt64IfZero(v any, def int64) int64 {
	value := NewValue(v).Int64Value()
	if value == 0 {
		return def
	}
	return value
}

func (n *Number) DefaultFloat64IfZero(v any, def float64) float64 {
	value := NewValue(v).Float64Value()
	if value == 0 {
		return def
	}
	return value
}
