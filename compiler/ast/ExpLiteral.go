package ast

import (
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
)

type Literal struct {
	token *token.Token
}

func NewLiteral(token *token.Token) *Literal {
	return &Literal{token}
}

func (l *Literal) expressionNode() {}
func (l *Literal) GetToken() *token.Token {
	return l.token
}

func (l *Literal) Env(env *environment.Environment) (any, error) {
	return l.token.Literal, nil
}
