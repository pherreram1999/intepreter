package parser

import "pahm/intepreter/inter/token"

func (p *Parser) primary() {
	switch p.PreaAnalisis.Tipo {
	case token.True:
		p.Match(token.True)
	case token.False:
		p.Match(token.False)
	case token.Null:
		p.Match(token.Null)
	case token.Number:
		p.Match(token.Number)
	case token.String:
		p.Match(token.String)
	case token.Identifier:
		p.Match(token.Identifier)
	case token.LeftParen:
		p.Match(token.LeftParen)
		p.expression()
		p.Match(token.RightParen)
	case token.Input:
		p.Match(token.Input)
		p.Match(token.LeftParen)
		p.expression()
		p.Match(token.RightParen)
	}
}
