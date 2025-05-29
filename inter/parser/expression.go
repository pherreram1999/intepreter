package parser

import "pahm/intepreter/inter/ast"

func (p *Parser) expression() ast.Expression {
	return p.assignment()
}
