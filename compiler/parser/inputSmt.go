package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) inputSmt() ast.Expression {
	p.Match(token.Input)
	p.Match(token.LeftParen)
	p.expression()
	p.Match(token.RightParen)
	p.Match(token.Semicolon)
	return nil
}
