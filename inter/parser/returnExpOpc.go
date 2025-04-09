package parser

import "pahm/intepreter/inter/token"

func (p *Parser) returnExpOpec() {
	if p.PreaAnalisis.Tipo != token.Semicolon {
		return // epsilon
	}
	p.expression()
}
