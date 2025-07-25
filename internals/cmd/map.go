package cmd

import (
	"fmt"
	"github.com/print-agent/bd-pokedexcli-v2/internals/cfg"
	"github.com/print-agent/bd-pokedexcli-v2/internals/pkg"
)

const URLLocation = "https://pokeapi.co/api/v2/location/"
const URLLocationArea = "https://pokeapi.co/api/v2/location-area/"

// TODO: return proper errors
func commandMap(config *cfg.Config) error {
	var url string

	if config.NextLocationURL != nil && *config.NextLocationURL != "" {
		url = *config.NextLocationURL
	} else {
		url = URLLocationArea
	}

	res, err := pkg.GetPokemonLocations(url)
	if err != nil {
		return fmt.Errorf("Failed to get locations: %w", err)
	}

	fmt.Println("\n📍 ▶️ Locations:")
	for _, location := range res.Results {
		fmt.Printf(" %s\n", location.Name)
	}

	config.PreviousLocationURL = res.Previous
	config.NextLocationURL = res.Next

	if res.Next != nil && *res.Next != "" {
		fmt.Printf("\n(Use 'map' again to see more locations)\n")
	} else {
		fmt.Printf("\n(No more locations available)\n")
	}

	return nil

}
func commandMapBack(config *cfg.Config) error {
	if config.PreviousLocationURL == nil && *config.PreviousLocationURL == "" {
		fmt.Printf("☠️ ▶️ No previous locations available\n")
		return nil
	}

	url := *config.PreviousLocationURL

	res, err := pkg.GetPokemonLocations(url)
	if err != nil {
		return fmt.Errorf("Failed to get locations: %w", err)
	}

	fmt.Println("\n📍 ▶️ Locations:")
	for _, location := range res.Results {
		fmt.Printf(" %s\n", location.Name)
	}

	config.PreviousLocationURL = res.Previous
	config.NextLocationURL = res.Next

	if config.PreviousLocationURL != nil && *config.PreviousLocationURL != "" {
		fmt.Printf("\n(Use 'map' to see more locations or 'mapb' to go back)\n")
	} else {
		fmt.Printf("\n(At the beginning. Use 'map' to see more locations)\n")
	}

	return nil

}
