package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) factorT(left ast.Expression) ast.Expression {
	switch p.PreaAnalisis.Tipo {
	case token.Slash, token.Star:
		operator := p.PreaAnalisis
		p.Match(operator.Tipo)
		right := p.factor()
		return ast.NewArithmetic(left, operator, right)
	default:
		return left // epsilon
	}
}
