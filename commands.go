package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
	}
}

func commandHelp() error {
	commands := getCommands()

	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, v := range commands {
		fmt.Printf("%v: %v", v.name, v.description)
		fmt.Println()
	}

	fmt.Println()

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil;
}
