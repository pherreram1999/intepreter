package parser

import "pahm/intepreter/inter/token"

func (p *Parser) forStmt() {
	p.Match(token.For)
	p.Match(token.LeftParen)
	p.forStmtInit()
	p.forStmtCond()
	p.forStmtInc()
	p.Match(token.RightParen)
	p.statement()
}
