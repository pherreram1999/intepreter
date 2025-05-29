package ast

import "pahm/intepreter/inter/token"

type Expression interface {
	expressionNode()
	GetToken() *token.Token
}
type Statement interface {
	statementNode()
}
