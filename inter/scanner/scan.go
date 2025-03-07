package scanner

import (
	"pahm/intepreter/inter/token"
	"pahm/intepreter/inter/validators"
)

func Scan(source string) []*token.Token {

	validators.Init()

	var tokens []*token.Token
	var lexama string
	var state uint

	sourceLen := len(source)

	var linea uint = 1

	for i := 0; i < sourceLen; i++ {
		character := rune(source[i])

		if character == '\n' {
			linea++
		}

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
					Linea:  linea,
				}
				tokens = append(tokens, t)
			} else {
				state = 3
				i-- // regresamos una pocision
				t := &token.Token{
					Tipo:   token.Greater,
					Lexema: lexama,
					Linea:  linea,
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
					Linea:  linea,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Less,
					Lexema: lexama,
					Linea:  linea,
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
					Linea:  linea,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Equal,
					Lexema: lexama,
					Linea:  linea,
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
					Linea:  linea,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Bang,
					Lexema: lexama,
					Linea:  linea,
				})
			}

			lexama = ""
			state = 0
		} else if state == 0 && validators.IsLetter(string(character)) {
			state = 13
			lexama += string(character)
		} else if state == 13 {
			if validators.IsLetterOrNumber(string(character)) {
				lexama += string(character)
				// state = 13
			} else {
				// validar palabra resevada
				lexama = ""
			}
		}
	}

	return tokens
}
