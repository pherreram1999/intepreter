package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) unary() ast.Expression {
	switch p.PreaAnalisis.Tipo {
	case token.Bang, token.Minus:
		operator := p.PreaAnalisis
		p.Match(operator.Tipo)
		right := p.unary()
		return ast.NewUnary(right, operator)
	default:
		p.call()
	}
}
