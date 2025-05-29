package ast

type Literal struct {
	Value interface{}
}

func NewLiteral(value interface{}) *Literal {
	return &Literal{value}
}

func (l *Literal) expressionNode() {}
