package ast

import "pahm/intepreter/compiler/environment"

type ExpressionStatement struct {
	Expression Expression
}

func NewExpressionStatement(expression Expression) *ExpressionStatement {
	return &ExpressionStatement{Expression: expression}
}
func (b *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) Env(env *environment.Environment) (any, error) {
	return es.Expression.Env(env)
}
