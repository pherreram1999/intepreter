package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) callT(callee ast.Expression) ast.Expression {
	if p.PreaAnalisis.Tipo == token.LeftParen {
		leftParen := p.PreaAnalisis
		p.Match(token.LeftParen)
		argumentos := p.arguments()
		p.Match(token.RightParen)
		return ast.NewCallFunction(callee, argumentos, leftParen)
	} else {
		return callee // epsilon
	}
}
