package parser

import "pahm/intepreter/inter/token"

func (p *Parser) varInit() {
	if p.PreaAnalisis.Tipo != token.Equal {
		return // esta vacio pues
	}
	p.Match(token.Equal)
	p.expression()
}
