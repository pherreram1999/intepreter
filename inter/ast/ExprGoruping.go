package ast

type Goruping struct {
	Expression Expression
}

func NewGoruping(expression Expression) *Goruping {
	return &Goruping{expression}
}

func (g *Goruping) expressionNode() {}
