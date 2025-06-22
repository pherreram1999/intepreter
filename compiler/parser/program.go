package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) Program() []ast.Statement {
	var statements []ast.Statement
	for p.PreaAnalisis.Tipo != token.EOF {
		statements = append(statements, p.Declaration())
	}
	return statements
}
