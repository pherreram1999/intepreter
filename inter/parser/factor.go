package parser

func (p *Parser) factor() {
	p.unary()
	p.factorT()
}
