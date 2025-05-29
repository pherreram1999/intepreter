package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) equalityT(left ast.Expression) ast.Expression {
	expected := p.PreaAnalisis.Tipo
	var operator *token.Token
	var right ast.Expression
	operator = p.PreaAnalisis
	switch expected {
	case token.BangEqual, token.EqualEqual:
		p.Match(operator.Tipo)
		right = p.equality()
	default:
		return left
	}
	return ast.NewRelational(left, operator, right)
}
