package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemonLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetPokemonLocations(URL string) (*pokemonLocations, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to make a request: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	locations := pokemonLocations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON: %w", err)
	}

	return &locations, nil

}
