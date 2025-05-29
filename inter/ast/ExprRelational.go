package ast

import "pahm/intepreter/inter/token"

type Relational struct {
	left     Expression
	Operator *token.Token
	right    Expression
}

func NewRelational(left Expression, operator *token.Token, right Expression) *Relational {
	return &Relational{left, operator, right}
}

func (r *Relational) expressionNode()        {}
func (r *Relational) GetToken() *token.Token { return r.Operator }
