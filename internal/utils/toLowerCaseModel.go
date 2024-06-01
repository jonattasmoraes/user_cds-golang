package utils

import (
	"reflect"
	"strings"
)

func ToLowerCaseExcept(model interface{}, exceptions ...string) {
	v := reflect.ValueOf(model).Elem()
	excMap := map[string]struct{}{}

	for _, exc := range exceptions {
		excMap[exc] = struct{}{}
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		if field.Kind() == reflect.String {
			if _, found := excMap[fieldType.Name]; !found {
				field.SetString(strings.ToLower(field.String()))
			}
		}
	}
}
