package main

import (
	"fmt"
	"github.com/print-agent/bd-pokedexcli-v2/internals"
)

func main() {
	// NOTE: Initialize Pokedex
	fmt.Println("Pokedex CLI ▶️ version 0.0.0 ")
	internals.REPL()
}
