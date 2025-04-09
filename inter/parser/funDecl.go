package parser

import "pahm/intepreter/inter/token"

func (p *Parser) funDecl() {
	p.Match(token.Fun)
	p.Match(token.Identifier)
	p.Match(token.LeftParen)
	p.parameters()
	p.Match(token.RightParen)
	p.block()
}
