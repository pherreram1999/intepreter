package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) logicAnd() ast.Expression {
	left := p.equality()
	return p.logicAndT(left)
}
