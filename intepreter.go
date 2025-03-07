package main

import (
	"fmt"
	"pahm/intepreter/inter/scanner"
)

func execute(code string) {
	tokens := scanner.Scan(code)
	for _, token := range tokens {
		fmt.Println("Token", token)
	}
}
