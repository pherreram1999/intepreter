package parser

import "pahm/intepreter/inter/token"

func (p *Parser) comparisonT() {
	switch p.PreaAnalisis.Tipo {
	case token.Greater:
		p.Match(token.Greater)
		p.comparison()
	case token.GreaterEqual:
		p.Match(token.GreaterEqual)
		p.comparison()
	case token.Less:
		p.Match(token.Less)
		p.comparison()
	case token.LessEqual:
		p.Match(token.LessEqual)
		p.comparison()
	default:
		return // epsilon
	}
}
