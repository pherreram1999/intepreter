package token

import (
	"fmt"
)

type Token struct {
	Tipo         TipoToken
	Lexema       string
	Literal      interface{}
	PrintLexema  bool
	PrintLiteral bool
	Linea        uint
}

func (t *Token) String() string {
	if t.Tipo == EOF {
		return fmt.Sprintf("<%s,lexema: %s>", t.Tipo, t.Lexema)
	}

	if t.PrintLexema && t.PrintLiteral {
		return fmt.Sprintf(`<%s, lexema: %s, literal: %v (%T),linea: %d>`, t.Tipo, t.Lexema, t.Literal, t.Literal, t.Linea)
	} else if t.PrintLexema {
		return fmt.Sprintf(`<%s, lexema: %s, linea: %d>`, t.Tipo, t.Lexema, t.Linea)
	} else if t.PrintLiteral && !t.PrintLexema {
		return fmt.Sprintf(`<%s,literal: %v (%T), linea: %d>`, t.Tipo, t.Literal, t.Literal, t.Linea)
	} else {
		return fmt.Sprintf(`<%s, linea: %d>`, t.Tipo, t.Linea)
	}
}
