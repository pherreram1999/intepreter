package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
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
