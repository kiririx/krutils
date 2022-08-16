package structx

import (
	"errors"
	"reflect"
)

// DeepCopy copy struct to another
func DeepCopy(src any, target any) error {
	if src == nil || target == nil {
		return errors.New("src or target is nil")
	}

	if reflect.TypeOf(target).Kind() != reflect.Pointer {
		return errors.New("target must be pointer")
	}

	var srcElem reflect.Value
	if reflect.TypeOf(src).Kind() == reflect.Pointer {
		srcElem = reflect.ValueOf(src).Elem()
	} else {
		srcElem = reflect.ValueOf(src)
	}

	// src is struct, target is struct
	if reflect.TypeOf(srcElem).Kind() == reflect.Struct && reflect.TypeOf(reflect.ValueOf(target).Elem()).Kind() == reflect.Struct {
		targetElem := reflect.ValueOf(target).Elem()
		for i := 0; i < srcElem.NumField(); i++ {
			val := srcElem.Field(i)
			name := srcElem.Type().Field(i).Name
			tVal := targetElem.FieldByName(name)
			if tVal.IsValid() && tVal.Kind() == val.Kind() {
				if val.Kind() == reflect.Struct && tVal.Kind() == reflect.Struct {
					_ = DeepCopy(val, tVal)
				}
				tVal.Set(val)
			}
		}
		return nil
	}
	return nil
}
