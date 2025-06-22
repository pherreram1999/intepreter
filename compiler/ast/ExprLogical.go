package ast

import (
	"fmt"
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
	"pahm/intepreter/compiler/validators"
)

type Logical struct {
	Left     Expression
	Operator *token.Token
	Right    Expression
}

func NewLogical(left Expression, operator *token.Token, right Expression) *Logical {
	return &Logical{left, operator, right}
}

func (l *Logical) expressionNode()        {}
func (l *Logical) GetToken() *token.Token { return l.Operator }

func (l *Logical) Env(env *environment.Environment) (any, error) {
	left, err := l.Left.Env(env)
	if err != nil {
		return nil, err
	}

	switch l.Operator.Tipo {
	case token.Or:
		if validators.IsTruthy(left) {
			return left, nil
		}
		return l.Right.Env(env)
	case token.And:
		if !validators.IsTruthy(left) {
			return left, nil
		}
		return l.Right.Env(env)
	}
	return nil, fmt.Errorf("operador lógico desconocido: %s", l.Operator.Lexema)
}
