package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) block() *ast.Block {
	p.Match(token.LeftBrace)
	var statements []ast.Statement
	for p.PreaAnalisis.Tipo != token.RightBrace && p.PreaAnalisis.Tipo != token.EOF {
		statements = append(statements, p.Declaration())
	}
	p.Match(token.RightBrace)
	return ast.NewBlock(statements)
}
