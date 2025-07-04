package validators

import (
	"pahm/intepreter/compiler/token"
)

func IsReservedWord(word string) token.TipoToken {
	switch word {
	case "else":
		return token.Else
	case "fun":
		return token.Fun
	case "print":
		return token.Print
	case "var":
		return token.Var
	case "false":
		return token.False
	case "if":
		return token.If
	case "return":
		return token.Return
	case "while":
		return token.While
	case "for":
		return token.For
	case "null":
		return token.Null
	case "true":
		return token.True
	case "input":
		return token.Input
	case "or":
		return token.Or
	case "and":
		return token.And
	default:
		return ""
	}
}
