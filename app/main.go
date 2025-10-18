package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		splittedInput := strings.Split(userInput, " ")
		clearedInput := clearInput(splittedInput)

		command, exists := commands.Commands[strings.ToLower(clearedInput[0])]

		if !exists {
			fmt.Printf("%v: command not found\n", clearedInput[0])
			continue
		}
		config := commands.Config{
			Args: clearedInput[1:],
		}
		error := command.Callback(&config)

		if error != nil {
			fmt.Printf("Error occurs: %v\n", error.Error())
		}

	}
}

func clearInput(args []string) (cleared []string) {
	for _, element := range args {
		cleared = append(cleared, strings.TrimSpace(element))
	}
	return
}
