package command

import (
	"fmt"
	"os"

	"github.com/groveshell/grove-shell/internal/env"
)

type PWDCommand struct{}

func (pwd PWDCommand) Name() string {
	return "pwd"
}

func (pwd PWDCommand) Description() string {
	return "print the current working directory"
}

func (pwd PWDCommand) Usage() string {
	return "pwd"
}

func (pwd PWDCommand) Run(args []string, env env.ShellEnvironment) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(wd)

	return nil
}
