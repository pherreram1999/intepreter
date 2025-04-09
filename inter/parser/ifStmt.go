package parser

import "pahm/intepreter/inter/token"

func (p *Parser) ifStmt() {
	p.Match(token.If)
	p.Match(token.LeftParen)
	p.expression()
	p.Match(token.RightParen)
	p.statement()
	p.elseStmt()
}
