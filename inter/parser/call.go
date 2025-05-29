package parser

import "pahm/intepreter/inter/ast"

func (p *Parser) call() ast.Expression {
	callee := p.primary()
	return p.callT(callee)
}
