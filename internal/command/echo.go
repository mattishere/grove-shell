package command

import (
	"fmt"

	"github.com/mattishere/grove-shell/internal/env"
	"github.com/mattishere/grove-shell/internal/utils"
)

type EchoCommand struct{}

func (echo EchoCommand) Name() string {
	return "echo"
}

func (echo EchoCommand) Description() string {
    return "print text and variables in the terminal"
}

func (echo EchoCommand) Usage() string {
    return "echo [text/\"text\"]..."
}

func (echo EchoCommand) Run(args []string, env env.ShellEnvironment) error {
	var msg string

	for _, arg := range args {
		if utils.IsString(arg) || utils.IsRawString(arg) {
			msg += arg[1:len(arg)-1] + " "
		} else {
			msg += arg + " "
		}
	}

	fmt.Println(msg)

	return nil
}

