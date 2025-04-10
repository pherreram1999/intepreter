package parser

import (
	"pahm/intepreter/inter/token"
)

func (p *Parser) inputSmt() {
	p.Match(token.Input)
	p.Match(token.LeftParen)
	p.expression()
	p.Match(token.RightParen)
	p.Match(token.Semicolon)
}
