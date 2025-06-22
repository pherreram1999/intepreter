package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) expStmt() ast.Statement {
	expr := p.expression()
	p.Match(token.Semicolon)
	return ast.NewExpressionStatement(expr)
}
