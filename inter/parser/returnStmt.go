package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) returnStmt() ast.Statement {
	p.Match(token.Return)
	value := p.returnExpOpec()
	p.Match(token.Semicolon)
	return ast.NewReturnStatement(value)
}
