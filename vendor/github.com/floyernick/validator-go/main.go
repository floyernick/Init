package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"unicode"

	"github.com/google/uuid"
)

type validationKeys struct {
	Pattern  *regexp.Regexp
	Min      int
	Max      int
	Len      int
	Digits   bool
	Letters  bool
	Required bool
	UUID     bool
	Datetime bool
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
			var keys validationKeys
			validationString := scheme.Field(i).Tag.Get("validate")
			if len(validationString) != 0 {
				validationList := strings.Split(validationString, "; ")
				for i := 0; i < len(validationList); i++ {
					validationKey := strings.SplitAfterN(validationList[i], "=", 2)
					switch validationKey[0] {
					case "required":
						keys.Required = true
					case "digits":
						keys.Digits = true
					case "letters":
						keys.Letters = true
					case "uuid":
						keys.UUID = true
					case "datetime":
						keys.Datetime = true
					}
					validationKey[0] = validationKey[0][:len(validationKey[0])-1]
					switch validationKey[0] {
					case "pattern":
						keys.Pattern, _ = regexp.Compile(validationKey[1])
					case "min":
						keys.Min, _ = strconv.Atoi(validationKey[1])
					case "max":
						keys.Max, _ = strconv.Atoi(validationKey[1])
					case "len":
						keys.Len, _ = strconv.Atoi(validationKey[1])
					}
				}
			}
			if !traverse(field, keys) {
				return false
			}
		}
	case reflect.Slice:
		if !keys.Required && data.Len() == 0 {
			return true
		}
		if keys.Min != 0 && data.Len() < keys.Min {
			return false
		}
		if keys.Max != 0 && data.Len() > keys.Max {
			return false
		}
		if keys.Len != 0 && data.Len() != keys.Len {
			return false
		}
		for u := 0; u < data.Len(); u++ {
			if !traverse(data.Index(u), keys) {
				return false
			}
		}
	case reflect.String:
		value := data.String()
		if !keys.Required && len(value) == 0 {
			return true
		}
		if keys.Letters {
			for _, r := range value {
				if !unicode.IsLetter(r) {
					return false
				}
			}
		}
		if keys.Digits {
			for _, r := range value {
				if !unicode.IsDigit(r) {
					return false
				}
			}
		}
		if keys.UUID {
			if _, err := uuid.Parse(value); err != nil {
				return false
			}
		}
		if keys.Datetime {
			if _, err := time.Parse(time.RFC3339, value); err != nil {
				return false
			}
		}
		if keys.Min != 0 && utf8.RuneCountInString(value) < keys.Min {
			return false
		}
		if keys.Max != 0 && utf8.RuneCountInString(value) > keys.Max {
			return false
		}
		if keys.Len != 0 && utf8.RuneCountInString(value) != keys.Len {
			return false
		}
		if keys.Pattern != nil && !keys.Pattern.MatchString(value) {
			return false
		}
	case reflect.Int:
		value := int(data.Int())
		if !keys.Required && value == 0 {
			return true
		}
		if keys.Min != 0 && value < keys.Min {
			return false
		}
		if keys.Max != 0 && value > keys.Max {
			return false
		}
		if keys.Len != 0 && value != keys.Len {
			return false
		}
	default:
		value := fmt.Sprintf("%v", data)
		if !keys.Required && len(value) == 0 {
			return true
		}
		if keys.Pattern != nil && !keys.Pattern.MatchString(value) {
			return false
		}
	}
	return true
}
