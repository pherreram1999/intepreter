package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) forStmtCond() ast.Expression {
	var condExpr ast.Expression
	if p.PreaAnalisis.Tipo == token.Semicolon {
		if isExpression(p.PreaAnalisis.Tipo) {
			condExpr = p.expression()
		} else {
			p.Error(token.Semicolon)
		}
	}
	p.Match(token.Semicolon)
	return condExpr
}
