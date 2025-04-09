package parser

import "pahm/intepreter/inter/token"

func (p *Parser) logicOrT() {
	if p.PreaAnalisis.Tipo != token.Or {
		return // epsilon
	}
	p.Match(token.Or)
	p.logicOr()
}
