package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) Program() []ast.Statement {
	var statements []ast.Statement
	for p.PreaAnalisis.Tipo != token.EOF {
		statements = append(statements, p.Declaration())
	}
	return statements
}
