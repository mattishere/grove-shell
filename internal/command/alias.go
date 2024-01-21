package command

import (
	"fmt"

	"github.com/groveshell/grove-shell/internal/env"
	"github.com/groveshell/grove-shell/internal/utils"
)

type AliasCommand struct{}

func (alias AliasCommand) Name() string {
	return "alias"
}

func (alias AliasCommand) Description() string {
    return "create a new alias"
}

func (alias AliasCommand) Usage() string {
    return "alias [name] [command/\"command\"]"
}

func (alias AliasCommand) Run(args []string, env env.ShellEnvironment) error {
	if len(args) == 2 {
		aliasKey := args[0]
		if utils.IsString(aliasKey) || utils.IsRawString(aliasKey) {
			return fmt.Errorf("invalid alias name (cannot be string)")
		}

		command := args[1]
		if utils.IsString(command) || utils.IsRawString(command) {
			command = command[1 : len(command)-1]
		}
		env.Aliases[aliasKey] = command
	} else if len(args) == 0 {
        final := "Aliases:\n"
		for name, alias := range env.Aliases {
			final += name + " -> " + alias + "\n"
		}
        fmt.Println(final)
	}

	return nil
}
