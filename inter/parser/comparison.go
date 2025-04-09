package parser

func (p *Parser) comparison() {
	p.term()
	p.comparisonT()
}
