package parser

func (p *Parser) logicOr() {
	p.logicAnd()
	p.logicOrT()
}
