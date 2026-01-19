package main

import (
	"fmt"
)

func commandExplore(cfg *config, location string) error {
	if location == "" {
		return fmt.Errorf("No location provided to explore")
	}

	fmt.Printf("Exploring %s...\n", location)

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(nil, location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range pokemonResp.Encounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}
