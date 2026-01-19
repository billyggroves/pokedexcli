package main

import (
	"fmt"
)

func commandListPokedex(cfg *config, ignore string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("No pokemon in your pokedex!")
		return nil
	}

	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
