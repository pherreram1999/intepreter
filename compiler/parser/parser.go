package parser

import (
	"fmt"
	"log"
	"pahm/intepreter/compiler/token"
)

type Parser struct {
	Tokens       []*token.Token
	Index        int
	PreaAnalisis *token.Token
}

func NewParser(tokens []*token.Token) *Parser {
	return &Parser{
		Tokens:       tokens,
		Index:        0,
		PreaAnalisis: tokens[0],
	}
}

func (p *Parser) Error(expected token.TipoToken) {
	if p.PreaAnalisis.Tipo == token.String {
		log.Printf(fmt.Sprintf(
			`Error sintáctico en la línea %d. Se esperaba %s pero se recibió "%s"`,
			p.PreaAnalisis.Linea,
			expected,
			p.PreaAnalisis.Lexema,
		))
	} else {
		log.Println(fmt.Sprintf(
			"Error sintáctico en la línea %d. Se esperaba %s pero se recibió %s",
			p.PreaAnalisis.Linea,
			expected,
			p.PreaAnalisis.Lexema,
		))
	}
}

func (p *Parser) Match(expected token.TipoToken) {
	if p.PreaAnalisis.Tipo == expected {
		// avanzamos al siguiente tipo
		p.Index++
		if p.Index < len(p.Tokens) {
			p.PreaAnalisis = p.Tokens[p.Index]
		}
	} else {
		p.Error(expected)
	}
}
