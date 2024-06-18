package main

import (
	"fmt"
)

func commandHelp(_ *config) error {
	fmt.Print("Welcome to the Pokedex!")
	fmt.Print("\n\n")
	fmt.Print("Usage:")
	fmt.Print("\n\n")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%v: %v \n", command.name, command.description)
	}

	return nil
}
