package parser

import (
	"pahm/intepreter/inter/token"
)

func (p *Parser) printStmt() {
	p.Match(token.Print)
	p.expression()
	p.Match(token.Semicolon)
}
