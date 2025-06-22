package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

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
	} else if isExpression(p.PreaAnalisis.Tipo) {
		return p.expStmt()
	} else {
		p.Error("Error sinstatico, se espera un statement, se recibio: " + p.PreaAnalisis.Tipo)
		return nil
	}
}
