package parser

func (p *Parser) forStmtInc() {
	if isExpression(p.PreaAnalisis.Tipo) {
		p.expression()
	}
}
