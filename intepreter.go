package main

import (
	"fmt"
	"log"
	"pahm/intepreter/inter/parser"
	"pahm/intepreter/inter/scanner"
	"pahm/intepreter/inter/token"
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
	// ejecutamos el programa
	p := parser.NewParser(tokens)
	p.Program()
	//printTokens(tokens)

	// si esta esta correcto debio llegar al EOF
	if p.PreaAnalisis.Tipo == token.EOF {
		fmt.Println("Todo well")
	} else {
		log.Printf("todo wrong: token actual | %s\n", p.PreaAnalisis)
	}

}
