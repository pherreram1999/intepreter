package ast

import "pahm/intepreter/inter/token"

type VarStatement struct {
	Name         *token.Token
	Initializier Expression
}

func NewVarStatement(name *token.Token, initializier Expression) *VarStatement {
	return &VarStatement{name, initializier}
}

func (v *VarStatement) statementNode() {}
