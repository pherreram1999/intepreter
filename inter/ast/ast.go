package ast

type Expression interface {
	expressionNode()
}
type Statement interface {
	statementNode()
}
