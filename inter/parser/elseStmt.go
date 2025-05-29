package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) elseStmt() ast.Statement {
	if p.PreaAnalisis.Tipo != token.Else {
		return nil // epsilon
	}
	p.Match(token.Else)
	return p.statement()
}
