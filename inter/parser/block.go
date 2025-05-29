package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
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
