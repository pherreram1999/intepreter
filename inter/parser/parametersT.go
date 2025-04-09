package parser

import "pahm/intepreter/inter/token"

func (p *Parser) parametersT() {
	if p.PreaAnalisis.Tipo != token.Comma {
		return
	}
	p.Match(token.Comma)
	p.Match(token.Identifier)
	p.parametersT()
}
