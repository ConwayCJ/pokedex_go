package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokecache "github.com/conwaycj/pokedex_go/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	next    string
	prev    string
	cache   pokecache.Cache
	pokedex map[string]Pokemon
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
		"explore": {
			name:        "explore <location-name>",
			description: "Explores location name for pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect stats of a caught pokemon in our pokedex",
			callback:    commandInspect,
		},
	}
}

func cleanInput(s string) []string {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	return words
}

func start(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		// wait for user input
		scanner.Scan()
		// get what user text + sanitize input
		userInput := cleanInput(scanner.Text())

		commandName := userInput[0]
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}

		// get command from list of commands
		command, exists := getCommands()[commandName]

		// if command exists...
		if exists {
			err := command.callback(cfg, args...)

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
