package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Print("\nThanks for using the Pokedex. Goodbye!\n")
	os.Exit(0)
	return nil
}
