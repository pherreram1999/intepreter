package token

import "fmt"

type Token struct {
	Tipo        TipoToken
	Lexema      string
	Literal     interface{}
	PrintLexema bool
	Linea       uint
}

func (t *Token) String() string {
	if t.PrintLexema {
		return fmt.Sprintf(`<%s, lexema: %s, linea: %d>`, t.Tipo, t.Lexema, t.Linea)
	}
	return fmt.Sprintf(`<%s, linea: %d>`, t.Tipo, t.Linea)
}
