package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func sayGooBay() {
	fmt.Println("Good Bye!")
	os.Exit(0)
}

func repl() {
	fmt.Println("*** Modo REPL ***")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			sayGooBay()
		}

		switch input {
		case "exit":
			sayGooBay()
		}
		execute(input)
	}
}
