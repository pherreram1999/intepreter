package parser

import "pahm/intepreter/inter/token"

func (p *Parser) forStmtInc() {
	if p.PreaAnalisis.Tipo == token.RightParen {
		return // epsilon
	} else {
		p.expression()
	}
}
