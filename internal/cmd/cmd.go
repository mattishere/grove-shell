package cmd

import (
	"os"
)

type Command interface {
	Name() string
	Run([]string) error
}

type CdCommand struct{}

func (cd CdCommand) Name() string {
	return "cd"
}

// to-do: add support for spaces! ("/projects/grove shell" e.g.)
func (cd CdCommand) Run(args []string) error {
	if len(args) == 0 || args[0] == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		os.Chdir(home)
	} else {
		os.Chdir(args[0])
	}

	return nil
}
