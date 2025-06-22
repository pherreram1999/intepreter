package ast

import "pahm/intepreter/compiler/token"

type FunctionStatement struct {
	Name       *token.Token
	Parameters []*token.Token
	Body       *Block
}

func NewFunctionStatement(token *token.Token, parameters []*token.Token, body *Block) *FunctionStatement {
	return &FunctionStatement{Name: token, Parameters: parameters, Body: body}
}
func (b *FunctionStatement) statementNode() {}
