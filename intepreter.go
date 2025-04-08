package main

import (
	"log"
	"pahm/intepreter/inter/parser"
	"pahm/intepreter/inter/scanner"
)

func execute(code string) {
	tokens := scanner.Scan(code)
	if len(tokens) == 0 {
		log.Fatalln("no token found")
	}
	// ejecutamos el programa

	p := parser.NewParser(tokens)
}
