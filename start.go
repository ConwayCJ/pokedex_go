package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next string
	prev string
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Repeat command for next 20.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world. Repeat command for previous 20.",
			callback:    commandMapb,
		},
	}
}

func sanitizeText(s string) string {
	s = strings.ToLower(s)
	return s
}

func start(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		// wait for user input
		scanner.Scan()
		// get what user typed
		userInput := scanner.Text()
		// user to lowercase
		userInput = sanitizeText(userInput)

		// get command from list of commands
		command, exists := getCommands()[userInput]

		// if command exists...
		if exists {
			err := command.callback(cfg)

			// if there's an error...
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}
