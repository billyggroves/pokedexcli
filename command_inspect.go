package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("No pokemon provided to explore")
	}
	if len(cfg.pokedex) == 0 {
		fmt.Printf("%s is not in the pokedex, you need to catch it!\n", pokemon)
		return nil
	}

	pm, ok := cfg.pokedex[pokemon]
	if !ok {
		fmt.Printf("%s is not in the pokedex, you need to catch it!\n", pokemon)
		return nil
	}

	// fmt.Println(pm)

	fmt.Printf("Name: %s\n", pm.Name)
	fmt.Printf("Height: %d\n", pm.Height)
	fmt.Printf("Weight: %d\n", pm.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pm.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pm.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	// fmt.Printf("Name: %s\n", pm.Name)
	// fmt.Printf("Name: %s\n", pm.Name)
	return nil
}
