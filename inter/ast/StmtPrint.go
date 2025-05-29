package ast

type Print struct {
	Expression Expression
}

func NewPrint(expression Expression) *Print {
	return &Print{Expression: expression}
}
func (p *Print) statementNode() {}
