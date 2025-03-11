package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSpace(input)

		switch input {
		case "exit":
			sayGooBay()
		}
		execute(input)
	}
}
