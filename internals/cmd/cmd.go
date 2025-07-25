package cmd

import (
	"fmt"
	"github.com/print-agent/bd-pokedexcli-v2/internals/cfg"
)

type pokeCommand struct {
	name        string
	description string
	callback    func(*cfg.Config) error
}

func getCommands() map[string]pokeCommand {
	return map[string]pokeCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"quit": {
			name:        "quit",
			description: "Quit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show names of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show names of the previous 20 locations",
			callback:    commandMapBack,
		},
	}
}

func PokemonCommands(name string, config *cfg.Config) error {
	commands := getCommands()
	command, exists := commands[name]
	if !exists {
		fmt.Printf("☠️ ▶️ Command is not found\n")
		return nil
	}

	return command.callback(config)
}
