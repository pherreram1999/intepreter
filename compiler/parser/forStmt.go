package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
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
