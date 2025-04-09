package parser

import "pahm/intepreter/inter/token"

func (p *Parser) callT() {
	if p.PreaAnalisis.Tipo != token.LeftParen {
		return // epsilon
	}
	p.Match(token.LeftParen)
	p.arguments()
	p.Match(token.RightParen)
}
