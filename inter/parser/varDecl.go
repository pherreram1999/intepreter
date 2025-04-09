package parser

import "pahm/intepreter/inter/token"

func (p *Parser) varDecl() {
	p.Match(token.Var)
	p.Match(token.Identifier)
	p.varInit()
	p.Match(token.Semicolon)
}
