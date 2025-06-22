package ast

import (
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
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

func (a *Assign) Env(env *environment.Environment) (any, error) {
	value, err := a.Value.Env(env)
	if err != nil {
		return nil, err
	}
	if err = env.Assing(a.Name, value); err != nil {
		return nil, err
	}
	return value, nil
}
