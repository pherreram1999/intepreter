package parser

import "pahm/intepreter/inter/token"

func (p *Parser) callT() {
	if p.PreaAnalisis.Tipo == token.LeftParen {

		p.Match(token.LeftParen)
		p.arguments()
		p.Match(token.RightParen)

	} else {
		return // epsilon
	}
}
