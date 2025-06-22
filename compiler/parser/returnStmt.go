package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) returnStmt() ast.Statement {
	p.Match(token.Return)
	value := p.returnExpOpec()
	p.Match(token.Semicolon)
	return ast.NewReturnStatement(value)
}
