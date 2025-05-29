package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) whileStmt() ast.Statement {
	p.Match(token.While)
	p.Match(token.LeftParen)
	condition := p.expression()
	p.Match(token.RightParen)
	body := p.statement()
	return ast.NewWhileStatement(condition, body)
}
