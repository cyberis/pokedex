package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *cliConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
