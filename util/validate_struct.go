package util

import (
	"errors"
	"fmt"
	"reflect"
)

func ValidateStruct(s interface{}) (err error) {
	structType := reflect.TypeOf(s)
	kind := structType.Kind()
	if kind != reflect.Struct {
		return errors.New("input param should be a struct")
	}

	structVal := reflect.ValueOf(s)
	fieldNum := structVal.NumField()

	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name

		isSet := field.IsValid() && !field.IsZero()

		// Set to true because bool is always set (i.e. defaults to true)
		if field.Type().Kind() == reflect.Bool {
			isSet = true
		}

		if !isSet {
			return fmt.Errorf("%s is not set", fieldName)
		}
	}

	return err
}
