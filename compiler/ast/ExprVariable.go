package ast

import (
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
)

type Variable struct {
	Token *token.Token
}

func NewVariable(token *token.Token) *Variable {
	return &Variable{token}
}
func (v *Variable) expressionNode() {}
func (v *Variable) GetToken() *token.Token {
	return v.Token
}
func (v *Variable) Env(env *environment.Environment) (any, error) {
	return env.Get(v.Token)
}
