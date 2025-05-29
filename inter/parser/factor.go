package parser

import "pahm/intepreter/inter/ast"

func (p *Parser) factor() ast.Expression {
	left := p.unary()
	return p.factorT(left)
}
