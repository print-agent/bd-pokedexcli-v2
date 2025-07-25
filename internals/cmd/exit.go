package cmd

import (
	"fmt"
	"github.com/print-agent/bd-pokedexcli-v2/internals/cfg"
	"os"
)

func commandExit(config *cfg.Config) error {
	fmt.Printf("🤢 ▶️ Closing the Pokedex... Goodbye! 🫰")
	os.Exit(0)
	return nil
}
