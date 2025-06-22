package validators

import "reflect"

// IsNumberV valida por reflexion si la variable es un tipo numerico
func IsNumberV(value any) bool {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	}
	return false
}
