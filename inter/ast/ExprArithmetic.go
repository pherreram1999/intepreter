package ast

import "pahm/intepreter/inter/token"

type Arithmetic struct {
	left     Expression
	operator *token.Token
	right    Expression
}

func NewArithmetic(left Expression, operator *token.Token, right Expression) *Arithmetic {
	return &Arithmetic{left, operator, right}
}

func (a *Arithmetic) expressionNode() {}
