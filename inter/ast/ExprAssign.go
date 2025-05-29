package ast

import "pahm/intepreter/inter/token"

type Assign struct {
	Name  *token.Token
	Value Expression
}

func NewAssign(name *token.Token, value Expression) *Assign {
	return &Assign{name, value}
}

func (a *Assign) expressionNode() {}
