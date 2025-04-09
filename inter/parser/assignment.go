package parser

func (p *Parser) assignment() {
	p.logicOr()
	p.assignmentOpc()
}
