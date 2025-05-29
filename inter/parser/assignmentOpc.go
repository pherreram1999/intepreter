package parser

import (
	"fmt"
	"pahm/intepreter/inter/ast"
	"pahm/intepreter/inter/token"
)

func (p *Parser) assignmentOpc(left ast.Expression) ast.Expression {
	if p.PreaAnalisis.Tipo != token.Equal {
		return left
	}
	assignToken := p.PreaAnalisis
	p.Match(token.Equal)
	value := p.expression()

	if variable, ok := value.(*ast.Variable); ok {
		return ast.NewAssign(variable.GetToken(), value)
	} else {
		fmt.Sprintf("Se esperaba un token '=' en la linea %d \n", assignToken.Linea)
		p.Error(token.Identifier)
		return left
	}
}
