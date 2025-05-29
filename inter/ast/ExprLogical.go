package ast

import "pahm/intepreter/inter/token"

type Logical struct {
	Left     Expression
	Operator *token.Token
	Right    Expression
}

func NewLogical(left Expression, operator *token.Token, right Expression) *Logical {
	return &Logical{left, operator, right}
}

func (l *Logical) expressionNode() {}
