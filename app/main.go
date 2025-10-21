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
		currentDir, err := os.Getwd()
		if err != nil {
			panic("Shell cannot work in current directory")
		}

		if !exists {
			config := commands.Config{
				Args:             clearedInput,
				CurrentDirectory: currentDir,
			}
			err := commands.ExecCommand(&config)
			if err != nil {
				fmt.Printf("%v", err)
			}
			continue
		}
		config := commands.Config{
			Args:             clearedInput[1:],
			CurrentDirectory: currentDir,
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
