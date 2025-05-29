package ast

import (
	"pahm/intepreter/inter/token"
)

type Assign struct {
	Name  *token.Token
	Value Expression
}

func NewAssign(name *token.Token, value Expression) Expression {
	return &Assign{name, value}
}

func (a *Assign) expressionNode()        {}
func (a *Assign) GetToken() *token.Token { return a.Name }
