package parser

func (p *Parser) term() {
	p.factor()
	p.termT()
}
