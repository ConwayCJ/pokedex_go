package main

import (
	"fmt"
	"os"
)

func commandExit(_ *config, _ ...string) error {
	fmt.Print("\nThanks for using the Pokedex. Goodbye!\n")
	os.Exit(0)
	return nil
}
