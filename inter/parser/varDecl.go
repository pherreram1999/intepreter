package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) varDecl() ast.Statement {
	p.Match(token.Var)
	name := p.PreaAnalisis
	p.Match(token.Identifier)
	initializer := p.varInit()
	p.Match(token.Semicolon)
	return ast.NewVarStatement(name, initializer)
}
