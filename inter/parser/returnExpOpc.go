package parser

func (p *Parser) returnExpOpec() {
	/*if p.PreaAnalisis.Tipo != token.Semicolon {
	//	return // epsilon
	}
	*/
	if isExpression(p.PreaAnalisis.Tipo) {
		p.expression()
	}
}
