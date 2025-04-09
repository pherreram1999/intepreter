package parser

func (p *Parser) call() {
	p.primary()
	p.callT()
}
