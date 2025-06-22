package ast

import (
	"errors"
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
	"pahm/intepreter/compiler/validators"
)

type Unary struct {
	Operator *token.Token
	Right    Expression
}

func NewUnary(Right Expression, operator *token.Token) *Unary {
	return &Unary{operator, Right}
}

func (u *Unary) expressionNode()        {}
func (u *Unary) GetToken() *token.Token { return u.Operator }

func (u *Unary) Env(env *environment.Environment) (any, error) {
	rigth, err := u.Right.Env(env)

	if err != nil {
		return nil, err
	}

	switch u.Operator.Tipo {
	case token.Minus:
		if validators.IsNumberV(rigth) {
			return rigth.(float64), nil // por convecion es flotante 64
		}
		return nil, errors.New("operando incompatible para -")
	case token.Bang:
		return !validators.IsTruthy(rigth), nil
	}
	return nil, errors.New("operador unario desconocido")
}
