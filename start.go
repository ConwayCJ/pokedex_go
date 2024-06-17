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
	callback    func() error
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
	}
}

func sanitizeText(s string) string {
	s = strings.ToLower(s)
	return s
}

func start() {
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
			err := command.callback()

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
