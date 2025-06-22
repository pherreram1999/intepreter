package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) whileStmt() ast.Statement {
	p.Match(token.While)
	p.Match(token.LeftParen)
	condition := p.expression()
	p.Match(token.RightParen)
	body := p.statement()
	return ast.NewWhileStatement(condition, body)
}
