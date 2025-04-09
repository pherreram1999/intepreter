package parser

import "pahm/intepreter/inter/token"

func (p *Parser) factorT() {
	switch p.PreaAnalisis.Tipo {
	case token.Slash:
		p.Match(token.Slash)
		p.factor()
	case token.Star:
		p.Match(token.Star)
		p.factor()
	default:
		return // epsilon

	}
}
