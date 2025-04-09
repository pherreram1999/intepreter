package parser

import "pahm/intepreter/inter/token"

func (p *Parser) forStmtCond() {
	if p.PreaAnalisis.Tipo == token.Semicolon {
		p.Match(token.Semicolon)
	} else {
		p.expression()
		p.Match(token.Semicolon)
	}
}
