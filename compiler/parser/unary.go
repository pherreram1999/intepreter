package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) unary() ast.Expression {
	switch p.PreaAnalisis.Tipo {
	case token.Bang, token.Minus:
		operator := p.PreaAnalisis
		p.Match(operator.Tipo)
		right := p.unary()
		return ast.NewUnary(right, operator)
	default:
		return p.call()
	}
}
