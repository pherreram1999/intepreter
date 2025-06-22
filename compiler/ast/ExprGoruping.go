package ast

import (
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
)

type Grouping struct {
	Expression Expression
}

func NewGrouping(expression Expression) *Grouping {
	return &Grouping{expression}
}

func (g *Grouping) expressionNode() {}
func (g *Grouping) GetToken() *token.Token {
	return nil
}

func (g *Grouping) Env(env *environment.Environment) (any, error) {
	return g.Expression.Env(env)
}
