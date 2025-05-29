package ast

import "pahm/intepreter/inter/token"

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
