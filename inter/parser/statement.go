package parser

import "pahm/intepreter/inter/token"

// TODO checar si input entra qui
func (p *Parser) statement() {
	if p.PreaAnalisis.Tipo == token.LeftBrace {
		// es bloque
		p.block()
	} else if p.PreaAnalisis.Tipo == token.For {
		p.forStmt()
	} else if p.PreaAnalisis.Tipo == token.If {
		p.ifStmt()
	} else if p.PreaAnalisis.Tipo == token.Print {
		p.printStmt()
	} else if p.PreaAnalisis.Tipo == token.Return {
		p.returnStmt()
	} else if p.PreaAnalisis.Tipo == token.While {
		p.whileStmt()
	} else {
		p.expStmt()
	}
}
