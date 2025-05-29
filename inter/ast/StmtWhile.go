package ast

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

func NewWhileStatement(condition Expression, body Statement) *WhileStatement {
	return &WhileStatement{Condition: condition, Body: body}
}

func (w *WhileStatement) statementNode() {}
