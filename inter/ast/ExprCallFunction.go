package ast

import "pahm/intepreter/inter/token"

type CallFunction struct {
	Callee    Expression
	Arguments []Expression
	Paren     *token.Token
}

func NewCallFunction(callee Expression, arguments []Expression, paren *token.Token) *CallFunction {
	return &CallFunction{callee, arguments, paren}
}

func (cf *CallFunction) expressionNode()        {}
func (cf *CallFunction) GetToken() *token.Token { return cf.Paren }
