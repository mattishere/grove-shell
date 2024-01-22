package run

import (
	"os"
	"os/exec"

	"github.com/mattishere/grove-shell/internal/command"
	"github.com/mattishere/grove-shell/internal/env"
	"github.com/mattishere/grove-shell/internal/expand"
	"github.com/mattishere/grove-shell/internal/lex"
	"github.com/mattishere/grove-shell/internal/utils"
)

func RunCommand(handler *command.CommandHandler, input string, shellEnv env.ShellEnvironment) error {
	tokens := lex.Lex(input)
	if len(tokens) == 0 {
		return nil
	}
	commandName := tokens[0]
	args := tokens[1:]

	expanderHandler := expand.NewExpanderHandler(
		&expand.HomeDirExpander{},
		&expand.EnvironmentExpander{},
	)

	for i, arg := range args {
		if !utils.IsRawString(args[i]) {
			args[i] = expanderHandler.Expand(arg)
		}
	}

	if value, ok := shellEnv.Aliases[commandName]; ok {
		aliasTokens := lex.Lex(value)
		commandName = aliasTokens[0]
		args = append(aliasTokens[1:], args...)
	}

	exists, err := handler.RunCmd(commandName, args, shellEnv)
	if err != nil {
		return err
	}

	if !exists {
		command := exec.Command(commandName, args...)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
            return err
		}
	}

	return nil
}
