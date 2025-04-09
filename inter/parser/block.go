package parser

import "pahm/intepreter/inter/token"

func (p *Parser) block() {
	p.Match(token.LeftBrace)
	p.Declaration()
	p.Match(token.RightBrace)
}
