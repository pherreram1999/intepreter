package validators

import (
	"fmt"
	"regexp"
	"strconv"
)

func ValidarLiteralNumerico(num string) interface{} {
	isfloat, _ := regexp.MatchString(".", num)
	isexp, _ := regexp.MatchString("E", num)
	exp := regexp.MustCompile("E*")
	fmt.Println("Si encontro un E")
	if isexp {
		expval := exp.FindString(num)
		fmt.Println(expval)
		num = exp.ReplaceAllString(num, "")

	}
	if !isfloat {
		strconv.Atoi(num)
	}
	return nil
}
