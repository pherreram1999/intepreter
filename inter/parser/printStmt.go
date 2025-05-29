package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) printStmt() ast.Statement {
	p.Match(token.Print)
	expr := p.expression()
	p.Match(token.Semicolon)
	return ast.NewPrint(expr)
}
