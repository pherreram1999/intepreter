package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func isStatement(t token.TipoToken) bool {
	if t == token.Print ||
		t == token.If ||
		t == token.For ||
		t == token.Return ||
		t == token.LeftBrace || // para los bloques
		t == token.While {
		return true
	}
	return false
}

func isExpression(t token.TipoToken) bool {
	if t == token.Identifier ||
		t == token.True ||
		t == token.False ||
		t == token.Null ||
		t == token.Number ||
		t == token.String ||
		t == token.Input ||
		t == token.LeftParen {
		return true
	}
	return false
}

func (p *Parser) Declaration() ast.Statement {
	if p.PreaAnalisis.Tipo == token.Fun {
		return p.funDecl()
	} else if p.PreaAnalisis.Tipo == token.Var {
		return p.varDecl()
	} else if isStatement(p.PreaAnalisis.Tipo) || isExpression(p.PreaAnalisis.Tipo) {
		return p.statement()
	}
	p.Error("error sintactico, se espera una declaracion")
	return nil
}
