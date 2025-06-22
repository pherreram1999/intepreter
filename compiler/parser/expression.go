package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) expression() ast.Expression {
	return p.assignment()
}
