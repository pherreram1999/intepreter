package validators

import (
	token2 "pahm/intepreter/inter/token"
)

func IsReservedWord(word string) token2.TipoToken {
	switch word {
	case "else":
		return token2.Else
	case "fun":
		return token2.Fun
	case "print":
		return token2.Print
	case "var":
		return token2.Var
	case "false":
		return token2.False
	case "if":
		return token2.If
	case "return":
		return token2.Return
	case "while":
		return token2.While
	case "for":
		return token2.For
	case "null":
		return token2.Null
	case "true":
		return token2.True
	default:
		return ""
	}
}
