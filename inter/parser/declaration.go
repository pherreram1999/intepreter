package parser

import "pahm/intepreter/inter/token"

func (p *Parser) Declaration() {
	switch p.PreaAnalisis.Tipo {
	case token.Fun:
		p.funDecl()
		p.Declaration()
	case token.Var:
		p.varDecl()
		p.Declaration()
	case token.Print:
		p.statement()
		p.Declaration()
	default:
		return // epsilon (ε) (vacío)
	}
}
