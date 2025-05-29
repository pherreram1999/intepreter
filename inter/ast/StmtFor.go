package ast

import "pahm/intepreter/inter/token"

type ForStatement struct {
	Initializer Statement
	Condition   Statement
	Increment   Statement
	Body        Statement
}

func NewForStatement(initializer Statement, condition, increment, body Statement) *ForStatement {
	return &ForStatement{initializer, condition, increment, body}
}
func (p *ForStatement) statementNode() {}
func (p *ForStatement) GetToken() *token.Token {
	return nil
}
