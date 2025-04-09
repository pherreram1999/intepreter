package parser

import "pahm/intepreter/inter/token"

func (p *Parser) expStmt() {
	p.expression()
	p.Match(token.Semicolon)
}
