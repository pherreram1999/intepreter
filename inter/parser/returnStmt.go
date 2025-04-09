package parser

import "pahm/intepreter/inter/token"

func (p *Parser) returnStmt() {
	p.Match(token.Return)
	p.returnExpOpec()
	p.Match(token.Semicolon)
}
