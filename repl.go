package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin) //	Scanner that reads the value typed by the user

	for {
		fmt.Print(" >")

		scanner.Scan()         //	Scans what was previously typed
		text := scanner.Text() //	Gets the text typed in the scanner

		cleaned := cleanInput(text) //	Cleans the input
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0] //	Defines that the command is the first word of the "cleaned" slice

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.callback()

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string { //	Function that cleans the input -> lowers the input
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
