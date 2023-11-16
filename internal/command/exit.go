package command

import (
	"os"
	"strconv"

	"github.com/groveshell/grove-shell/internal/env"
)

type ExitCommand struct{}

func (exit ExitCommand) Name() string {
	return "exit"
}

func (exit ExitCommand) Description() string {
	return "exit the shell"
}

func (exit ExitCommand) Usage() string {
	return "exit [code (0-255)]"
}

func (exit ExitCommand) Run(args []string, env env.ShellEnvironment) error {
	if len(args) == 0 {
		os.Exit(0)
	} else {
		exitCode, err := strconv.Atoi(args[0])
		if err != nil || exitCode < 0 || exitCode > 255 {
			os.Exit(0)
		} else {
			os.Exit(exitCode)
		}
	}

	return nil
}
