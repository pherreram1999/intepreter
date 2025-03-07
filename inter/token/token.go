package token

import "fmt"

type Token struct {
	Tipo    TipoToken
	Lexema  string
	Literal interface{}
	Linea   uint
}

func (t *Token) String() string {
	return fmt.Sprintf(`<%s>`, t.Tipo)
}
