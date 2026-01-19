package main

import (
	"fmt"
	"math/rand"

	"github.com/billyggroves/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("No Pokemon provided to catch")
	}
	if cfg.pokedex == nil {
		cfg.pokedex = make(map[string]pokeapi.Pokemon)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	pokemonResp, err := cfg.pokeapiClient.GetPokemonInfo(nil, pokemon)
	if err != nil {
		return err
	}

	chance := rand.Intn(100)

	switch {
	case pokemonResp.BaseExp > 300:
		if 100 >= chance && chance > 95 {
			cfg.pokedex[pokemon] = pokemonResp
			fmt.Printf("%s was caught!\n", pokemon)
		} else {
			fmt.Printf("%s escaped!\n", pokemon)
		}
	case pokemonResp.BaseExp > 200:
		if 100 >= chance && chance > 85 {
			cfg.pokedex[pokemon] = pokemonResp
			fmt.Printf("%s was caught!\n", pokemon)
		} else {
			fmt.Printf("%s escaped!\n", pokemon)
		}
	case pokemonResp.BaseExp > 100:
		if 100 >= chance && chance > 70 {
			cfg.pokedex[pokemon] = pokemonResp
			fmt.Printf("%s was caught!\n", pokemon)
		} else {
			fmt.Printf("%s escaped!\n", pokemon)
		}
	default:
		if 100 >= chance && chance > 50 {
			cfg.pokedex[pokemon] = pokemonResp
			fmt.Printf("%s was caught!\n", pokemon)
		} else {
			fmt.Printf("%s escaped!\n", pokemon)
		}
	}

	return nil
}
