package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide 1 pokemon name to inspect")
	}
	reqPoke := args[0]

	pokemon, exists := cfg.pokedex[reqPoke]
	if !exists {
		return errors.New("bleep bloop, you haven't caught this pokemon yet")
	}

	fmt.Printf("\nName: %s", pokemon.Name)
	fmt.Printf("ID %v: ", pokemon.ID)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, v := range pokemon.Stats {
		fmt.Printf("   -%s: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Type: ")
	for _, v := range pokemon.Types {
		fmt.Printf("   -%s\n", v.Type.Name)
	}
	return nil
}
