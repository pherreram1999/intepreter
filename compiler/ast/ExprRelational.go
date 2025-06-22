package ast

import (
	"errors"
	"fmt"
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
	"pahm/intepreter/compiler/validators"
)

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

func (r *Relational) Env(env *environment.Environment) (any, error) {
	left, err := r.left.Env(env)
	if err != nil {
		return nil, err
	}
	right, err := r.right.Env(env)
	if err != nil {
		return nil, err
	}

	if !validators.IsNumberV(left) {
		return nil, fmt.Errorf("El expresion relacion en la izquierda no es numerica")
	}

	if !validators.IsNumberV(right) {
		return nil, fmt.Errorf("El expresion relacion en la derecha no es numerica")
	}

	a := left.(float64)
	b := right.(float64)

	switch r.Operator.Tipo {
	case token.Greater:
		return a > b, nil
	case token.GreaterEqual:
		return a >= b, nil
	case token.Less:
		return a < b, nil
	case token.LessEqual:
		return a <= b, nil
	case token.EqualEqual:
		return a == b, nil
	case token.BangEqual:
		return a != b, nil
	}

	return nil, errors.New("No encontro un operador valido")
}
