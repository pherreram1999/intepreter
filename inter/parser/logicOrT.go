package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) logicOrT(left ast.Expression) ast.Expression {
	if p.PreaAnalisis.Tipo != token.Or {
		return left // epsilon
	}
	operator := p.PreaAnalisis
	p.Match(token.Or)
	right := p.logicOr()
	return ast.NewLogical(left, operator, right)
}
