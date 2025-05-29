package ast

type ReturnStatement struct {
	Value Expression
}

func NewReturnStatement(value Expression) *ReturnStatement {
	return &ReturnStatement{Value: value}
}
func (r *ReturnStatement) statementNode() {}
