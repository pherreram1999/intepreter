package ast

type IfStatement struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

func NewIfStatement(condition Expression, thenBranch Statement, elseBranch Statement) *IfStatement {
	return &IfStatement{Condition: nil, ThenBranch: nil, ElseBranch: nil}
}

func (i *IfStatement) statementNode() {}
