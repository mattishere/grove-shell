package run

import (
	"os"
	"os/exec"

	"github.com/groveshell/grove-shell/internal/cmd"
	"github.com/groveshell/grove-shell/internal/config"
	"github.com/groveshell/grove-shell/internal/expand"
	"github.com/groveshell/grove-shell/internal/lex"
	"github.com/groveshell/grove-shell/internal/utils"
)

func RunCommand(handler *cmd.CommandHandler, input string, config config.Config) error {
	tokens := lex.Lex(input)
	commandName := tokens[0]
	args := tokens[1:]

    aliases := config.Aliases

	expanderHandler := expand.NewExpanderHandler(
		&expand.HomeDirExpander{},
		&expand.EnvironmentExpander{},
	)

	for i, arg := range args {
		if !utils.IsRawString(args[i]) {
			args[i] = expanderHandler.Expand(arg)
		}
	}

    if value, ok := aliases[commandName]; ok {
        aliasTokens := lex.Lex(value)
        commandName = aliasTokens[0]
        args = append(aliasTokens[1:], args...)
    }

	exists, err := handler.RunCmd(commandName, args)
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
