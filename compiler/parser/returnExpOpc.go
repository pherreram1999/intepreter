package parser

import "pahm/intepreter/compiler/ast"

func (p *Parser) returnExpOpec() ast.Expression {
	if isExpression(p.PreaAnalisis.Tipo) {
		return p.expression()
	}
	return nil
}
