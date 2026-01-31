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
	callback    func(*cliConfig) error
}

type cliConfig struct {
	nextURL string
	prevURL string
}

func repl() {
	config := cliConfig{nextURL: "", prevURL: ""}

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
		if cmd, exists := getCommands()[command]; exists {
			err := cmd.callback(&config)
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
			callback:    commandMap,
		},
	}
	return commands
}
