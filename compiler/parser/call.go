package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) call() ast.Expression {
	callee := p.primary()
	return p.callT(callee)
}
