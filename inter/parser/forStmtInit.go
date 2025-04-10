package parser

import "pahm/intepreter/inter/token"

func (p *Parser) forStmtInit() {
	if p.PreaAnalisis.Tipo == token.Var {
		p.varDecl()
	} else if isExpression(p.PreaAnalisis.Tipo) {
		p.expStmt()
	} else {
		p.Match(token.Semicolon)
	}
}
