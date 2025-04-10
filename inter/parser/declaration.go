package parser

import (
	"pahm/intepreter/inter/token"
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

func (p *Parser) Declaration() {
	if p.PreaAnalisis.Tipo == token.Fun {
		p.funDecl()
		p.Declaration()
	} else if p.PreaAnalisis.Tipo == token.Var {
		p.varDecl()
		p.Declaration()
	} else if isStatement(p.PreaAnalisis.Tipo) || isExpression(p.PreaAnalisis.Tipo) {
		p.statement()
		p.Declaration()
	}

}
