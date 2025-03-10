package scanner

import (
	"pahm/intepreter/inter/token"
	"pahm/intepreter/inter/validators"
)

func Scan(source string) []*token.Token {

	validators.Init()

	var tokens []*token.Token
	var lexema string
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
			lexema += string(character)
		} else if state == 1 {
			lexema += string(character)
			if character == '=' {
				state = 2
				t := &token.Token{
					Tipo:   token.GreaterEqual,
					Lexema: lexema,
					Linea:  linea,
				}
				tokens = append(tokens, t)
			} else {
				state = 3
				i-- // regresamos una pocision
				t := &token.Token{
					Tipo:   token.Greater,
					Lexema: lexema,
					Linea:  linea,
				}
				tokens = append(tokens, t)
			}
			lexema = ""
			state = 0

		} else if state == 0 && character == '<' {
			state = 4
			lexema += string(character)
		} else if state == 4 { // menor o igual
			lexema += string(character)
			if character == '=' {
				tokens = append(tokens, &token.Token{
					Tipo:   token.LessEqual,
					Lexema: lexema,
					Linea:  linea,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Less,
					Lexema: lexema,
					Linea:  linea,
				})
			}
			state = 0
			lexema = ""
		} else if state == 0 && character == '=' {
			state = 7
			lexema += string(character)
		} else if state == 7 {
			lexema += string(character)
			if character == '=' {
				tokens = append(tokens, &token.Token{
					Tipo:   token.EqualEqual,
					Lexema: lexema,
					Linea:  linea,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Equal,
					Lexema: lexema,
					Linea:  linea,
				})
			}

			lexema = ""
			state = 0
		} else if state == 0 && character == '!' {
			state = 10
			lexema += string(character)
		} else if state == 10 {
			lexema += string(character)
			if character == '=' {
				tokens = append(tokens, &token.Token{
					Tipo:   token.BangEqual,
					Lexema: lexema,
					Linea:  linea,
				})
			} else {
				i--
				tokens = append(tokens, &token.Token{
					Tipo:   token.Bang,
					Lexema: lexema,
					Linea:  linea,
				})
			}

			lexema = ""
			state = 0
		} else if state == 0 && validators.IsLetter(string(character)) {
			state = 13
			lexema += string(character)
		} else if state == 13 {
			if validators.IsLetterOrNumber(string(character)) {
				lexema += string(character)
				// state = 13
			} else {
				// validar palabra resevada
				lexema = ""
			}
		}
	}

	return tokens
}
