package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) logicOr() ast.Expression {
	left := p.logicAnd()
	return p.logicOrT(left)
}
