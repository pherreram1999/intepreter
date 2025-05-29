package ast

import "pahm/intepreter/inter/token"

type InputExpression struct {
	Name  *token.Token
	Value Expression
}

func NewInputExpression(name *token.Token, value Expression) *InputExpression {
	return &InputExpression{name, value}
}
func (ie *InputExpression) expressionNode() {}
func (ie *InputExpression) GetToken() *token.Token {
	return ie.Name
}
