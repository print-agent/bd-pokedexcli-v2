package cmd

import (
	"fmt"
	"github.com/print-agent/bd-pokedexcli-v2/internals/cfg"
)

func commandHelp(config *cfg.Config) error {
	fmt.Println("🥶 ▶️ Welcome to the Pokedex!")
	fmt.Println("🥶 ▶️ Usage: ")

	for name, cmd := range getCommands() {
		fmt.Printf("🥶 ▶️ %s: %s\n", name, cmd.description)
	}

	return nil
}
