package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) equality() ast.Expression {
	left := p.comparison()
	return p.equalityT(left)

}
