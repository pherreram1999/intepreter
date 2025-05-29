package ast

import "pahm/intepreter/inter/token"

type Unary struct {
	Operator token.Token
	Right    Expression
}

func NewUnary(left Expression, operator token.Token, right Expression) *Unary {
	return &Unary{operator, left}
}

func (u *Unary) expressionNode() {}
