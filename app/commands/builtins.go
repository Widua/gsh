package commands

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

func TypeCommand(config *Config) error {
	if len(config.Args) < 1 {
		return fmt.Errorf("This command require at least one parameter")
	}
	cmdName := config.Args[0]

	_, exists := Commands[cmdName]

	if exists {

		fmt.Printf("%v is a shell builtin\n", cmdName)
		return nil
	}

	path := os.Getenv("PATH")

	splittedPath := strings.Split(path, ":")

	for _, pathenv := range splittedPath {
		foundPath := searchPath(pathenv, cmdName)
		if foundPath != "" {
			fmt.Printf("%v is %v\n", cmdName, foundPath)
			return nil
		}

	}

	fmt.Printf("%v: not found\n", cmdName)

	return nil
}

func searchPath(path string, cmdName string) string {
	found := ""
	walkfn := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}
		if info.Name() == cmdName && !info.IsDir() && strings.Contains(info.Mode().String(), "x") {
			found = path
			return filepath.SkipAll
		}
		return nil
	}
	filepath.Walk(path, walkfn)
	if len(found) > 0 {
		return found
	}
	return ""

}
