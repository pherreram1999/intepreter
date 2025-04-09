package parser

import "pahm/intepreter/inter/token"

func (p *Parser) parameters() {
	p.Match(token.Identifier)
	p.parametersT()
}
