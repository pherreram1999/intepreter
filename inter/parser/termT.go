package parser

import "pahm/intepreter/inter/token"

func (p *Parser) termT() {
	switch p.PreaAnalisis.Tipo {
	case token.Minus:
		p.Match(token.Minus)
		p.term()
	case token.Plus:
		p.Match(token.Plus)
		p.term()
	default:
		return // epsilon
	}
}
