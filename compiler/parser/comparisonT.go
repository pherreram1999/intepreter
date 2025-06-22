package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) comparisonT(left ast.Expression) ast.Expression {
	switch p.PreaAnalisis.Tipo {
	case token.Greater, token.GreaterEqual, token.Less, token.LessEqual:
		operator := p.PreaAnalisis
		p.Match(operator.Tipo)
		right := p.comparison()
		return ast.NewRelational(left, operator, right)
	default:
		return left // epsilon
	}
}
