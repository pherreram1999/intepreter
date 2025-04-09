package parser

import "pahm/intepreter/inter/token"

func (p *Parser) assignmentOpc() {
	if p.PreaAnalisis.Tipo != token.Equal {
		return
	}
	p.Match(token.Equal)
	p.expression()
}
