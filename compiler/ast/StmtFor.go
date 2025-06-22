package ast

import "pahm/intepreter/compiler/token"

type ForStatement struct {
	Initializer Statement
	Condition   Expression
	Increment   Expression
	Body        Statement
}

func NewForStatement(initializer Statement, condition Expression, increment Expression, body Statement) *ForStatement {
	return &ForStatement{initializer, condition, increment, body}
}
func (p *ForStatement) statementNode() {}
func (p *ForStatement) GetToken() *token.Token {
	return nil
}
