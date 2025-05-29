package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) termT(left ast.Expression) ast.Expression {
	switch p.PreaAnalisis.Tipo {
	case token.Minus, token.Plus:
		operator := p.PreaAnalisis
		p.Match(operator.Tipo)
		right := p.term()
		return ast.NewArithmetic(left, operator, right)
	default:
		return left // epsilon
	}
}
