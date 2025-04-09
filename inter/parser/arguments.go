package parser

import "pahm/intepreter/inter/token"

func (p *Parser) arguments() {
	if p.PreaAnalisis.Tipo == token.RightParen {
		// significa que llegamos al final de los argumentos, sin nada que hacer
		return // epsilon
	}
	p.expression()
	p.argumentsT()
}
