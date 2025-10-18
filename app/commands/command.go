package commands

type Command struct {
	Name     string
	Callback func(config *Config) error
}

type Config struct {
	Args []string
}

var Commands map[string]Command = map[string]Command{
	"exit": {
		Name:     "Exit",
		Callback: ExitCommand,
	},
}
