package validators

import (
	token2 "pahm/intepreter/inter/token"
)

func IsReservedWord(word string) (token2.TipoToken, bool) {
	switch word {
	case "else":
		return token2.Else, true
	case "fun":
		return token2.Fun, true
	case "print":
		return token2.Print, true
	case "var":
		return token2.Var, true
	case "false":
		return token2.False, true
	case "if":
		return token2.If, true
	case "return":
		return token2.Return, true
	case "while":
		return token2.While, true
	case "for":
		return token2.For, true
	case "null":
		return token2.Null, true
	case "true":
		return token2.True, true
	default:
		return "", false
	}
}
