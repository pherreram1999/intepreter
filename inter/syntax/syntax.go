package syntax

import "pahm/intepreter/inter/token"

func Analizador(tokens []*token.Token) {
	parser := newParser(tokens)

}
