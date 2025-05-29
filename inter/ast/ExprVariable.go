package ast

import "pahm/intepreter/inter/token"

type Variable struct {
	Name *token.Token
}

func NewVariable(name *token.Token) *Variable {
	return &Variable{name}
}
func (v *Variable) expressionNode() {}
