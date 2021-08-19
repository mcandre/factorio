package main

import (
	"os"

	"github.com/mcandre/factorio"
)

func main() {
	args := os.Args[1:]

	if err := factorio.Port(args); err != nil {
		panic(err)
	}
}
