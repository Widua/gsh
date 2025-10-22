package commands

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
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

	foundPath := searchPath(cmdName)

	if foundPath != "" {
		fmt.Printf("%v is %v\n", cmdName, foundPath)
		return nil

	}

	fmt.Printf("%v: not found\n", cmdName)

	return nil
}

func ExecCommand(config *Config) error {
	cmdName := config.Args[0]
	cmdArgs := config.Args[1:]

	cmdPath := searchPath(cmdName)

	if cmdPath == "" {
		return fmt.Errorf("%v: not found\n", cmdName)
	}

	cmd := exec.Command(cmdName, cmdArgs...)

	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func PwdCommand(config *Config) error {
	fmt.Println(config.CurrentDirectory)
	return nil
}

func CdCommand(config *Config) error {

	path := config.Args[0]

	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("cd: %v: No such file or directory", path)
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("cd: not a directory: %v", path)
	}
	os.Chdir(path)
	return nil
}

func searchPath(cmdName string) string {
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

	for _, path := range strings.Split(os.Getenv("PATH"), ":") {
		filepath.Walk(path, walkfn)
		if len(found) > 0 {
			return found
		}
	}

	return ""

}
