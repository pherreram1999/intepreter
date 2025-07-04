package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
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
