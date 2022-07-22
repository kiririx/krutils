package copyx

import (
	"errors"
	"reflect"
)

func DeepCopy(src any, target any) error {
	if src == nil || target == nil {
		return errors.New("src or target is nil")
	}
	if reflect.TypeOf(target).Kind() != reflect.Pointer {
		return errors.New("target must be pointer")
	}
	srcV := src
	if reflect.TypeOf(src).Kind() == reflect.Pointer {
		srcV = reflect.ValueOf(src).Elem()
	} else {
		srcV = &src
	}
	// src is struct, target is struct
	if reflect.TypeOf(srcV).Kind() == reflect.Struct && reflect.TypeOf(reflect.ValueOf(target).Elem()).Kind() == reflect.Struct {
		sElem := reflect.ValueOf(src).Elem()
		tElem := reflect.ValueOf(target).Elem()
		for i := 0; i < sElem.NumField(); i++ {
			val := sElem.Field(i)
			name := sElem.Type().Field(i).Name
			tVal := tElem.FieldByName(name)
			if tVal.IsValid() && tVal.Kind() == val.Kind() {
				if val.Kind() == reflect.Struct && tVal.Kind() == reflect.Struct {
					_ = DeepCopy(val, tVal)
				}
				tVal.Set(val)
			}
		}
		return nil
	}
	// src is struct, target is map
	// src is map, target is struct
	// src is map, target is map
	return nil
}
