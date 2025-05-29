package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) forStmt() ast.Statement {
	p.Match(token.For)
	p.Match(token.LeftParen)
	initializer := p.forStmtInit()
	condition := p.forStmtCond()
	increment := p.forStmtInc()
	p.Match(token.RightParen)
	body := p.statement()
	return ast.NewForStatement(initializer, condition, increment, body)
}
