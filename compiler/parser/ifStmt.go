package parser

import (
	"pahm/intepreter/compiler/ast"
	"pahm/intepreter/compiler/token"
)

func (p *Parser) ifStmt() ast.Statement {
	p.Match(token.If)
	p.Match(token.LeftParen)
	condition := p.expression()
	p.Match(token.RightParen)
	thenBranch := p.statement()
	var elseBranch ast.Statement
	elseBranch = p.elseStmt()
	return ast.NewIfStatement(condition, thenBranch, elseBranch)
}
