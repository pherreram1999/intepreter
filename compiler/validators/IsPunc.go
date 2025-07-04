package validators

import "pahm/intepreter/compiler/token"

func IsPunc(p rune) token.TipoToken {
	switch p {
	case '+':
		return token.Plus
	case '-':
		return token.Minus
	case '*':
		return token.Star
	case ';':
		return token.Semicolon
	case ',':
		return token.Comma
	case '(':
		return token.LeftParen
	case ')':
		return token.RightParen
	case '{':
		return token.LeftBrace
	case '}':
		return token.RightBrace
	default:
		return ""
	}
}
