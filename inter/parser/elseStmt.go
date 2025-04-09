package parser

import "pahm/intepreter/inter/token"

func (p *Parser) elseStmt() {
	if p.PreaAnalisis.Tipo != token.Else {
		return // epsilon
	}
	p.Match(token.Else)
	p.statement()
}
