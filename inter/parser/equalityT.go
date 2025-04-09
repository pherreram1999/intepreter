package parser

import "pahm/intepreter/inter/token"

// TODO revisar si debe lleva un bucle por lo de asociatividad por la izquierda
func (p *Parser) equalityT() {
	expected := p.PreaAnalisis.Tipo
	switch expected {
	case token.BangEqual:
		p.Match(token.BangEqual)
		p.equality()
	case token.EqualEqual:
		p.Match(token.EqualEqual)
		p.equality()
	default:
		return
	}
}
