package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cyberis/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig, ...string) error
}

type cliConfig struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
}

func repl(cfg *cliConfig) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		command := words[0]
		params := []string{}
		if len(words) > 1 {
			params = words[1:]
		}
		if cmd, exists := getCommands()[command]; exists {
			err := cmd.callback(cfg, params...)
			if err != nil {
				fmt.Printf("Error executing command %q: %v\n", command, err)
			}
		} else {
			fmt.Printf("Unknown command: %q\n", command)
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

// Here are our available commands
func getCommands() map[string]cliCommand {
	var commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the Pokedex map",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of the Pokedex map",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a specific Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a specific Pokemon in your Pokedex",
			callback:    commandInspect,
		},
	}
	return commands
}
