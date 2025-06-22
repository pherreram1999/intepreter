package ast

import (
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
)

type Node interface {
	Env(env *environment.Environment) (any, error)
}

type Expression interface {
	Node
	expressionNode()
	GetToken() *token.Token
}
type Statement interface {
	Node
	statementNode()
}
