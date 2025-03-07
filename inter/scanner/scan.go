package scanner

import (
	"pahm/intepreter/inter/token"
)

func Scan(source string) []*token.Token {
	var tokens []*token.Token
	var lexama string
	var state uint

	sourceLen := len(source)

	for i := 0; i < sourceLen; i++ {
		character := rune(source[i])

		if state == 0 && character == '>' {
			state = 1
			lexama += string(character)
		} else if state == 1 {
			lexama += string(character)
			if character == '=' {
				state = 2
				t := &token.Token{
					Tipo:   token.GreaterEqual,
					Lexema: lexama,
				}
				tokens = append(tokens, t)
			} else {
				state = 3
				i-- // regresamos una pocision
				t := &token.Token{
					Tipo:   token.Greater,
					Lexema: lexama,
				}
				tokens = append(tokens, t)
			}
			lexama = ""
			state = 0

		} else if state == 0 && character == '<' {
			state = 4
			lexama += string(character)
		} else if state == 4 { // menor o igual
			lexama += string(character)
			if character == '=' {
				tokens = append(tokens, &token.Token{
					Tipo:   token.LessEqual,
					Lexema: lexama,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Less,
					Lexema: lexama,
				})
			}
			state = 0
			lexama = ""
		} else if state == 0 && character == '=' {
			state = 7
			lexama += string(character)
		} else if state == 7 {
			lexama += string(character)
			if character == '=' {
				tokens = append(tokens, &token.Token{
					Tipo:   token.EqualEqual,
					Lexema: lexama,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Equal,
					Lexema: lexama,
				})
			}

			lexama = ""
			state = 0
		} else if state == 0 && character == '!' {
			state = 10
			lexama += string(character)
		} else if state == 10 {
			lexama += string(character)
			if character == '=' {
				tokens = append(tokens, &token.Token{
					Tipo:   token.BangEqual,
					Lexema: lexama,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Bang,
					Lexema: lexama,
				})
			}

			lexama = ""
			state = 0
		}
	}

	return tokens
}
