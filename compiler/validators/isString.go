package validators

import "reflect"

func IsString(value any) bool {
	r := reflect.ValueOf(value)
	return r.Kind() == reflect.String
}
