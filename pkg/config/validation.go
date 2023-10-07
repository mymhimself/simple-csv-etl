package config

import (
	"fmt"
	"reflect"
)

func ValidateStruct(s interface{}) error {
	return validateStructValue(reflect.ValueOf(s))
}

func validateStructValue(value reflect.Value) error {

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldTag := value.Type().Field(i).Tag.Get("validate")

		// Skip fields with "validate:"-" tag
		if fieldTag == "-" {
			continue
		}

		switch field.Kind() {
		case reflect.Ptr, reflect.Interface:
			// Check if the pointer or interface is nil
			if field.IsNil() {
				fieldName := value.Type().Field(i).Name
				return fmt.Errorf("%s is not initialized", fieldName)
			}
			if field.Kind() == reflect.Ptr {
				if fieldTag == "shallow" {
					continue
				}

				err := validateStructValue(field.Elem())
				if err != nil {
					return err
				}
			}
		case reflect.Struct:
			err := validateStructValue(field)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
