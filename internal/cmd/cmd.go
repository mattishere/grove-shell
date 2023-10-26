package cmd

import (
	"fmt"
	"os"

	"github.com/groveshell/grove-shell/internal/utils"
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

type EchoCommand struct{}

func (echo EchoCommand) Name() string {
    return "echo"
}

func (echo EchoCommand) Run(args []string) error {
    var msg string

    for _, arg := range args {
        if utils.IsString(arg) {
            // TO-DO: later, these will ignore insertions from the shell (~ -> home dir etc.)
            msg += arg[1:len(arg)-1] + " "
        } else {
            msg += arg + " "
        }
    }

    fmt.Println(msg)

    return nil
}
