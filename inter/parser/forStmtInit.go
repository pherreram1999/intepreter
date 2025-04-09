package parser

import "pahm/intepreter/inter/token"

func (p *Parser) forStmtInit() {
	if p.PreaAnalisis.Tipo == token.Var {
		p.varDecl()
	} else if p.PreaAnalisis.Tipo == token.Semicolon {
		p.Match(token.Semicolon)
	} else {
		p.expStmt()
	}
}
