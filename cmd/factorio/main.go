package main

import (
	"log"
	"os"

	"github.com/mcandre/factorio"
)

func main() {
	if len(os.Args) == 0 {
		log.Fatalf("error: missing program name\n")
	}

	args := os.Args[1:]

	if err := factorio.Port(args); err != nil {
		panic(err)
	}
}
