package parser

func (p *Parser) equality() {
	p.comparison()
	p.equalityT()
}
