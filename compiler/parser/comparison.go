package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) comparison() ast.Expression {
	left := p.term()
	return p.comparisonT(left)
}
