package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) forStmtInit() ast.Statement {
	if p.PreaAnalisis.Tipo == token.Var {
		return p.varDecl()
	} else if isExpression(p.PreaAnalisis.Tipo) {
		return p.expStmt()
	} else if p.PreaAnalisis.Tipo == token.Semicolon {
		p.Match(token.Semicolon)
		return nil
	} else {
		p.Error(token.Semicolon)
		return nil
	}
}
