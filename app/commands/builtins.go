package commands

import (
	"fmt"
	"os"
	"strconv"
)

func ExitCommand(config *Config) error {
	args := config.Args
	if len(args) > 1 {
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
