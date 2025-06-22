package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) funDecl() ast.Statement {
	p.Match(token.Fun)
	name := p.PreaAnalisis
	p.Match(token.Identifier)
	p.Match(token.LeftParen)
	parameters := p.parameters()
	p.Match(token.RightParen)
	body := p.block()
	return ast.NewFunctionStatement(name, parameters, body)
}
