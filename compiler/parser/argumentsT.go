package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) argumentsT() []ast.Expression {
	if p.PreaAnalisis.Tipo != token.Comma {
		return []ast.Expression{} // epsilon
	}
	var argumentos []ast.Expression
	firstArg := append(argumentos, p.expression())
	others := p.argumentsT()
	argumentos = append(firstArg, others...)
	return argumentos
}
