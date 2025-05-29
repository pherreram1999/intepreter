package parser

import "pahm/intepreter/inter/ast"

func (p *Parser) assignment() ast.Expression {

	left := p.logicOr()
	return p.assignmentOpc(left)
}
