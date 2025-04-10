package parser

import "pahm/intepreter/inter/token"

func (p *Parser) forStmtCond() {
	if isExpression(p.PreaAnalisis.Tipo) {
		p.expression()
		p.Match(token.Semicolon)
	} else {
		p.Match(token.Semicolon)
	}
}
