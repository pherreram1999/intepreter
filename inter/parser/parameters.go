package parser

import "pahm/intepreter/inter/token"

func (p *Parser) parameters() []*token.Token {
	var parameters []*token.Token
	if p.PreaAnalisis.Tipo == token.Identifier {
		parameters = append(parameters, p.PreaAnalisis)
		p.Match(token.Identifier)
		parameters = append(parameters, p.parametersT()...)
	}
	return parameters
}
