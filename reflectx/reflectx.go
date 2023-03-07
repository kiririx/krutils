package reflectx

import "reflect"

// RangeStructField 遍历所有的字段以及内嵌结构体
func RangeStructField(t reflect.Type, f func(field reflect.StructField)) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if t.Field(i).Anonymous {
			RangeStructField(field.Type, f)
		} else {
			f(field)
		}
	}
}
