package request

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type validationKeys struct {
	Pattern   *regexp.Regexp
	MinLength int
	MaxLength int
}

func Process(in interface{}) bool {
	data := reflect.ValueOf(in)
	return traverse(data, validationKeys{})
}

func traverse(data reflect.Value, keys validationKeys) bool {
	switch data.Kind() {
	case reflect.Struct:
		scheme := data.Type()
		for i := 0; i < data.NumField(); i++ {
			field := data.Field(i)
			keys := validationKeys{}
			if value := scheme.Field(i).Tag.Get("pattern"); value != "" {
				keys.Pattern, _ = regexp.Compile(value)
			}
			if value := scheme.Field(i).Tag.Get("min"); value != "" {
				keys.MinLength, _ = strconv.Atoi(value)

			}
			if value := scheme.Field(i).Tag.Get("max"); value != "" {
				keys.MaxLength, _ = strconv.Atoi(value)

			}
			if !traverse(field, keys) {
				return false
			}
		}
	case reflect.Slice:
		if keys.MinLength != 0 && data.Len() < keys.MinLength {
			return false
		}
		if keys.MaxLength != 0 && data.Len() > keys.MaxLength {
			return false
		}
		for u := 0; u < data.Len(); u++ {
			if !traverse(data.Index(u), keys) {
				return false
			}
		}
	default:
		value := fmt.Sprintf("%v", data)
		if keys.Pattern != nil && !keys.Pattern.MatchString(value) {
			return false
		}
	}
	return true
}
