package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("you haven't caught any pokemon yet")
	}

	fmt.Println("Your Pokedex: ")
	for k := range cfg.pokedex {
		fmt.Println("  - " + k)
	}

	return nil
}
