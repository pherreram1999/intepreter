package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) primary() ast.Expression {
	var tmp *token.Token
	switch p.PreaAnalisis.Tipo {
	case token.True:
		tmp = p.PreaAnalisis
		p.Match(token.True)
		return ast.NewLiteral(tmp)
	case token.False:
		tmp = p.PreaAnalisis
		p.Match(token.False)
		return ast.NewLiteral(tmp)
	case token.Null:
		tmp = p.PreaAnalisis
		p.Match(token.Null)
		return ast.NewLiteral(tmp)
	case token.Number:
		tmp = p.PreaAnalisis
		p.Match(token.Number)
		return ast.NewLiteral(tmp)
	case token.String:
		tmp = p.PreaAnalisis
		p.Match(token.String)
		return ast.NewLiteral(tmp)
	case token.Identifier:
		tmp = p.PreaAnalisis
		p.Match(token.Identifier)
		return ast.NewVariable(tmp)
	case token.LeftParen:
		p.Match(token.LeftParen)
		expr := p.expression()
		p.Match(token.RightParen)
		return ast.NewGrouping(expr)
	case token.Input:
		tmp = p.PreaAnalisis
		p.Match(token.Input)
		p.Match(token.LeftParen)
		var promptExpr ast.Expression
		if isExpression(p.PreaAnalisis.Tipo) {
			promptExpr = p.expression()
		} else if p.PreaAnalisis.Tipo != token.RightParen {
			// esta vacio y se cierra
			p.Error(token.RightParen)
		}
		p.Match(token.RightParen)
		return ast.NewInputExpression(tmp, promptExpr)
	default:
		p.Error(p.PreaAnalisis.Tipo)
		return nil
	}
}
