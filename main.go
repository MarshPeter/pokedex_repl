package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		text, err := reader.ReadString('\n')

		if err != nil {
			panic("Error reading from console: " + err.Error())
		}

		text = strings.Trim(text, "\n")

		if command, ok := commands[text]; ok {
			err := command.callback()

			if err != nil {
				panic(err.Error())
			}
		}
	}
}
