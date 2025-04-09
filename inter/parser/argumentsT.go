package parser

import "pahm/intepreter/inter/token"

func (p *Parser) argumentsT() {
	if p.PreaAnalisis.Tipo != token.Comma {
		return // epsilon
	}
	p.expression()
	p.argumentsT()
}
