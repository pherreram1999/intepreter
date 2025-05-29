package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) expStmt() ast.Statement {
	expr := p.expression()
	p.Match(token.Semicolon)
	return ast.NewExpressionStatement(expr)
}
