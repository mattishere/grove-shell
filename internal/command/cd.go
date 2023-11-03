package command

import (
	"fmt"
	"os"

	"github.com/groveshell/grove-shell/internal/env"
	"github.com/groveshell/grove-shell/internal/utils"
)

type CdCommand struct{}

func (cd CdCommand) Name() string {
	return "cd"
}

func (cd CdCommand) Description() string {
    return "change your current directory"
}

func (cd CdCommand) Usage() string {
    return "cd [path]"
}

func (cd CdCommand) Run(args []string, env env.ShellEnvironment) error {
	if len(args) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		os.Chdir(home)
		return nil
	}

	path := args[0]

	if utils.IsString(path) || utils.IsRawString(path) {
		path = path[1 : len(path)-1]
	}

	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("path error: %s not found", path)
	}

	return nil
}

