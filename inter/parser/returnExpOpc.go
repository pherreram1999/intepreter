package parser

func (p *Parser) returnExpOpec() {
	if isExpression(p.PreaAnalisis.Tipo) {
		p.expression()
	}
}
