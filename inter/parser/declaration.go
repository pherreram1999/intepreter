package parser

import "pahm/intepreter/inter/token"

func (p *Parser) Declaration() {
	switch p.PreaAnalisis.Tipo {
	case token.Fun:
		p.funDecl()
	case token.Var:
		p.varDecl()
	case token.Print:
		p.printStmt()
	case token.If:
		p.statement()
	case token.While:
		p.statement()
	case token.For:
		p.statement()
	}
}
