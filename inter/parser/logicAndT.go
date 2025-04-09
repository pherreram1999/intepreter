package parser

import "pahm/intepreter/inter/token"

func (p *Parser) logicAndT() {
	if p.PreaAnalisis.Tipo != token.And {
		return // epsilon
	}
	p.Match(token.And)
	p.logicAnd()
}
