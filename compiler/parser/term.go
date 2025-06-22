package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) term() ast.Expression {
	left := p.factor()
	return p.termT(left)

}
