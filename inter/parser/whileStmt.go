package parser

import "pahm/intepreter/inter/token"

func (p *Parser) whileStmt() {
	p.Match(token.While)
	p.Match(token.LeftParen)
	p.expression()
	p.Match(token.RightParen)
	p.statement()

}
