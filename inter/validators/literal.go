package validators

import (
	"strconv"
	"strings"
)

func ValidarLiteralNumerico(num string) interface{} {
	isFloat := strings.Contains(num, ".")
	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		panic(err)
	}
	if !isFloat {
		return int(f)
	}
	return f

}
