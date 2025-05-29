package ast

import "pahm/intepreter/inter/token"

type Unary struct {
	Operator *token.Token
	Right    Expression
}

func NewUnary(Right Expression, operator *token.Token) *Unary {
	return &Unary{operator, Right}
}

func (u *Unary) expressionNode()        {}
func (u *Unary) GetToken() *token.Token { return u.Operator }
