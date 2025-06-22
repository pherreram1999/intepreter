package ast

type ExpressionStatement struct {
	Expression Expression
}

func NewExpressionStatement(expression Expression) *ExpressionStatement {
	return &ExpressionStatement{Expression: expression}
}
func (b *ExpressionStatement) statementNode() {}
