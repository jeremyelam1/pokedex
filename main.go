package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Commands string

const (
	exit Commands = "exit"
	help Commands = "help"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	cliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Provides a help message",
			callback:    commandHelp,
		},
	}

	var command string
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		command = cleanInput(input.Text())[0]

		switch command {
		case string(exit):
			if err := cliCommands["exit"].callback(); err != nil {
				fmt.Println("Error:", err)
			}

		case string(help):
			if err := cliCommands["help"].callback(); err != nil {
				fmt.Println("Error:", err)
			}

		}

	}
}

func cleanInput(text string) []string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	text = reg.ReplaceAllString(strings.ToLower(text), "")

	splitText := strings.Split(text, " ")

	result := make([]string, 0, len(splitText))
	for _, value := range splitText {
		result = append(result, value)
	}

	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
		`)

	return nil
}
