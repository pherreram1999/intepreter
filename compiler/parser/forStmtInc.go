package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) forStmtInc() ast.Expression {
	if p.PreaAnalisis.Tipo != token.RightParen {
		if isExpression(p.PreaAnalisis.Tipo) {
			return p.expression()
		} else {
			p.Error(token.RightParen)
			return nil
		}
	}
	return nil
}
