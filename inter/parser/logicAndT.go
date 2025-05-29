package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) logicAndT(left ast.Expression) ast.Expression {
	if p.PreaAnalisis.Tipo != token.And {
		return left // epsilon
	}
	operator := p.PreaAnalisis
	p.Match(token.And)
	right := p.logicAnd()
	return ast.NewLogical(left, operator, right)
}
