package parser

import "pahm/intepreter/inter/token"

func Analizador(tokens []*token.Token) {
	parser := NewParser(tokens)

}
