package ast

import "pahm/intepreter/inter/token"

type Literal struct {
	token *token.Token
}

func NewLiteral(token *token.Token) *Literal {
	return &Literal{token}
}

func (l *Literal) expressionNode() {}
func (l *Literal) GetToken() *token.Token {
	return l.token
}
