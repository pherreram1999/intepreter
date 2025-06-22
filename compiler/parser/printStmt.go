package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) printStmt() ast.Statement {
	p.Match(token.Print)
	expr := p.expression()
	p.Match(token.Semicolon)
	return ast.NewPrint(expr)
}
