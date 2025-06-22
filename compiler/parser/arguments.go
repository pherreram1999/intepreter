package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) arguments() []ast.Expression {
	if p.PreaAnalisis.Tipo == token.RightParen {
		// significa que llegamos al final de los argumentos, sin nada que hacer
		return []ast.Expression{} // epsilon
	}
	var argumentos []ast.Expression
	argumentos = append(argumentos, p.expression())
	others := p.argumentsT()
	argumentos = append(argumentos, others...)
	return argumentos
}
