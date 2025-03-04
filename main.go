package main

import (
	"os"
)

func main() {
	args := os.Args[1:len(os.Args)]
	noArgs := len(args)
	if noArgs == 0 {
		// se lanza el modo REPL
		repl()
	} else if noArgs == 1 {
		// se espera sea un directorio que se encarga de extraer su contenido
		path := args[0]
		fileContent, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		fileBody := string(fileContent)
		execute(fileBody)
	} else {
		panic("Solo puedes especificar como maximo un path de entrada o sin argumentos para REPL")
	}

}
