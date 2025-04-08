package parser

import "pahm/intepreter/inter/token"

func (p *Parser) Declaration() {
	switch p.PreaAnalisis.Tipo {
	case token.Fun:
		p.funDecl()

	}
}
