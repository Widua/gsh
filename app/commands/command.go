package commands

type Command struct {
	Name     string
	Callback func(config *Config) error
}

type Config struct {
	Args []string
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"exit": {
			Name:     "Exit",
			Callback: ExitCommand,
		},
		"echo": {
			Name:     "Echo",
			Callback: EchoCommand,
		},
		"type": {
			Name:     "Type",
			Callback: TypeCommand,
		},
	}

}
