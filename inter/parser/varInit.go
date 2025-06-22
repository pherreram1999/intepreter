package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) varInit() ast.Expression {
	if p.PreaAnalisis.Tipo != token.Equal {
		p.Error("Se esperaba var init " + p.PreaAnalisis.Tipo)
		return nil // esta vacio pues
	}
	p.Match(token.Equal)
	return p.expression()
}
