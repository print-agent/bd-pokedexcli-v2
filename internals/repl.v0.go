package internals

import (
	"bufio"
	"fmt"
	"github.com/print-agent/bd-pokedexcli-v2/internals/cfg"
	"github.com/print-agent/bd-pokedexcli-v2/internals/cmd"
	"github.com/print-agent/bd-pokedexcli-v2/internals/pkg"
	"os"
)

func REPL() {
	config := &cfg.Config{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("😼 ▶️ ")

		if !scanner.Scan() {
			break
		}

		cmdText := pkg.CleanInput(scanner.Text())

		if len(cmdText) == 0 {
			continue
		}

		// TODO: checker if text matched defined commands

		err := cmd.PokemonCommands(cmdText[0], config)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
