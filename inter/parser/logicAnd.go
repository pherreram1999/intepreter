package parser

func (p *Parser) logicAnd() {
	p.equality()
	p.logicAndT()
}
