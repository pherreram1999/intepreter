package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) assignment() ast.Expression {
	left := p.logicOr()
	return p.assignmentOpc(left)
}
