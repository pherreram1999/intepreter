package scanner

import (
	"pahm/intepreter/inter/tokens"
)

func Scann(source string) []tokens.Token {
	var tokens []tokens.Token
	var lexama string
	var state uint

	sourceLen := len(source)

	for i := 0; i < sourceLen; i++ {
		character := rune(source[i])

		if state == 0 && character == '>' {
			state = 1
			lexama += string(character)
		}

	}

	return tokens
}
