package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) varInit() ast.Expression {
	if p.PreaAnalisis.Tipo != token.Equal {
		return nil // esta vacio pues
	}
	p.Match(token.Equal)
	return p.expression()
}
