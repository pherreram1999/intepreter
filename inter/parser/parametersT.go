package parser

import "pahm/intepreter/inter/token"

func (p *Parser) parametersT() []*token.Token {
	var parameters []*token.Token
	if p.PreaAnalisis.Tipo != token.Comma {
		return parameters
	}
	p.Match(token.Comma)
	if p.PreaAnalisis.Tipo == token.Identifier {
		parameters = append(parameters, p.PreaAnalisis)
		p.Match(token.Identifier)
		parameters = append(parameters, p.parametersT()...)
	} else {
		p.Error(token.Identifier)
	}
	return parameters
}
