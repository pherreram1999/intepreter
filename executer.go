package main

import (
	"fmt"
	"log"
	"pahm/intepreter/compiler/parser"
	"pahm/intepreter/compiler/scanner"
	"pahm/intepreter/compiler/token"
)

func printTokens(tokens []*token.Token) {
	for _, tok := range tokens {
		fmt.Println(tok)
	}
}

func execute(code string) {
	tokens := scanner.Scan(code)
	if len(tokens) == 0 {
		log.Fatalln("no token found")
	}
	p := parser.NewParser(tokens)
	statements := p.Program()

	// si esta esta correcto debio llegar al EOF
	if p.PreaAnalisis.Tipo == token.EOF {
		fmt.Println("programa valido")
	} else {
		log.Fatalln("!!!!! PROGRAMA INVALIDA :O  -> token actual", p.PreaAnalisis)
	}

}
