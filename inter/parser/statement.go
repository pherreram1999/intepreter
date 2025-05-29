package parser

import (
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

// TODO checar si input entra qui
func (p *Parser) statement() ast.Statement {
	if p.PreaAnalisis.Tipo == token.LeftBrace {
		// es bloque
		return p.block()
	} else if p.PreaAnalisis.Tipo == token.For {
		return p.forStmt()
	} else if p.PreaAnalisis.Tipo == token.If {
		return p.ifStmt()
	} else if p.PreaAnalisis.Tipo == token.Print {
		return p.printStmt()
	} else if p.PreaAnalisis.Tipo == token.Return {
		return p.returnStmt()
	} else if p.PreaAnalisis.Tipo == token.While {
		return p.whileStmt()
	} else if p.PreaAnalisis.Tipo == token.Input {
		return p.inputSmt() // la nueva caracteristica
	} else if isExpression(p.PreaAnalisis.Tipo) {
		return p.expStmt()
	} else {
		return nil
	}
}
