package scanner

import (
	"fmt"
	"pahm/intepreter/inter/token"
	"pahm/intepreter/inter/validators"
	"unicode"
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
		} else if state == 0 && unicode.IsLetter(character) {
			state = 13
			lexema += string(character)
		} else if state == 13 {
			if unicode.IsDigit(character) || unicode.IsLetter(character) {
				lexema += string(character)
				state = 13
			} else {
				// validar palabra resevada
				if TipoToken := validators.IsReservedWord(lexema); len(TipoToken) > 0 {
					TokenReservado := &token.Token{
						Tipo:   TipoToken,
						Lexema: lexema,
						Linea:  linea,
					}
					if TokenReservado.Tipo == token.True {
						TokenReservado.Literal = 1
					} else if TokenReservado.Tipo == token.False {
						TokenReservado.Literal = 0
					} else if TokenReservado.Tipo == token.Null {
						TokenReservado.Literal = nil
					}
					tokens = append(tokens, TokenReservado)

				} else { // si no es reservada es un identificador
					tokens = append(tokens, &token.Token{
						Tipo:        token.Identifier,
						Lexema:      lexema,
						Linea:       linea,
						PrintLexema: true,
					})
				}
				lexema = ""
				state = 0
				i--
			}

		} else if state == 0 && unicode.IsDigit(character) {
			state = 15
			lexema += string(character)
		} else if state == 15 {
			if unicode.IsDigit(character) {
				lexema += string(character)
				state = 15
			} else if character == '.' {
				state = 16
				lexema += string(character)
			} else if character == 'E' {
				state = 18
				lexema += string(character)
			} else {
				tokens = append(tokens, &token.Token{
					Tipo:        token.Number,
					Lexema:      lexema,
					PrintLexema: true,
					Linea:       linea,
					Literal:     validators.ValidarLiteralNumerico(lexema),
				})

				lexema = ""
				i--
				state = 0
			}
		} else if state == 16 && unicode.IsDigit(character) {
			lexema += string(character)
			state = 17
		} else if state == 17 {
			if unicode.IsDigit(character) {
				lexema += string(character)
				state = 17
			} else if character == 'E' {
				state = 18
				lexema += string(character)
			} else {
				tokens = append(tokens, &token.Token{
					Tipo:        token.Number,
					Lexema:      lexema,
					PrintLexema: true,
					Linea:       linea,
					Literal:     lexema,
				})
				state = 0
				i--
				lexema = ""
			}
		} else if state == 18 && (character == '+' || character == '-') {
			state = 19
			lexema += string(character)
		} else if state == 18 && unicode.IsDigit(character) {
			state = 20
			lexema += string(character)
		} else if state == 19 && unicode.IsDigit(character) {
			state = 20
			lexema += string(character)
		} else if state == 20 {
			if unicode.IsDigit(character) {
				lexema += string(character)
				state = 20
			} else {
				tokens = append(tokens, &token.Token{
					Tipo:        token.Number,
					Lexema:      lexema,
					PrintLexema: true,
					Linea:       linea,
					Literal:     lexema,
				})
				i--
				lexema = ""
				state = 0
			}
		} else if state == 0 && character == '"' {
			state = 24
		} else if state == 24 {
			if character == '"' {

				tokens = append(tokens, &token.Token{
					Tipo:        token.String,
					Lexema:      lexema,
					PrintLexema: true,
					Linea:       linea,
					Literal:     lexema,
				})

				lexema = ""
				state = 0
			} else if character == '\n' {
				panic("error: salto de linea")
			} else {
				lexema += string(character)
			}
		} else if state == 0 && character == '/' {
			state = 26
			//lexema += string(character)
		} else if state == 26 {
			if character == '/' {
				state = 30
				//lexema += string(character)
			} else if character == '*' {
				state = 27
				//lexema += string(character)
			} else {
				tokens = append(tokens, &token.Token{
					Tipo:   token.Slash,
					Lexema: lexema,
					Linea:  linea,
				})
				lexema = ""
				state = 0
				i--
			}
		} else if state == 27 {
			if character == '*' {
				state = 28
				//lexema += string(character)
			} else {
				state = 27
				//lexema += string(character)
			}
			// de lo contrario se queda en 27
		} else if state == 28 {
			if character == '*' {
				// se queda en el mismo
				//lexema += string(character)
			} else if character == '/' {
				state = 0
				lexema = ""
			} else {
				state = 27
				//lexema += string(character)
			}
		} else if state == 30 {
			if character == '\n' {
				state = 0
				lexema += ""
			} else {
				//lexema += string(character)
			}
		} else if state == 0 && unicode.IsSpace(character) || validators.IsDelim(character) {
			lexema += string(character)
			state = 33
		} else if state == 33 {
			if unicode.IsSpace(character) || validators.IsDelim(character) {
				lexema += string(character)
			} else {
				state = 0
				lexema = ""
				i--
			}
		} else if TipoToken := validators.IsPunc(character); len(TipoToken) > 0 && state == 0 {
			tokens = append(tokens, &token.Token{
				Tipo:   TipoToken,
				Lexema: lexema,
				Linea:  linea,
			})
			state = 0
			lexema = ""
		} else {
			panic(fmt.Sprintf("error: Simbolo no valido: %s", string(character)))
		}
		//fmt.Printf("State %d\t,caracter %s\t,index:%d\n", state, string(character), i)
	}

	return tokens
}
