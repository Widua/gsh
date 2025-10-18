package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExitCommand(config *Config) error {
	args := config.Args
	if len(args) < 1 {
		os.Exit(0)
	}
	exitCode := args[0]

	code, error := strconv.Atoi(exitCode)

	if error != nil {
		return fmt.Errorf("Argument: %v should be an integer", exitCode)
	}
	os.Exit(code)

	return nil
}

func EchoCommand(config *Config) error {
	content := strings.Join(config.Args, " ")
	fmt.Println(content)

	return nil
}
