package parser

import "pahm/intepreter/inter/token"

func (p *Parser) unary() {
	switch p.PreaAnalisis.Tipo {
	case token.Bang:
		p.Match(token.Bang)
		p.unary()
	case token.Minus:
		p.Match(token.Minus)
		p.unary()
	default:
		p.call()
	}
}
