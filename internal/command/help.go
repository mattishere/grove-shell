package command

import (
	"fmt"

	"github.com/groveshell/grove-shell/internal/env"
)

type HelpCommand struct{}

func (help HelpCommand) Name() string {
    return "help"
}

func (help HelpCommand) Description() string {
    return "get basic help and information about the shell"
}

func (help HelpCommand) Usage() string {
    return "help"
}

func (help HelpCommand) Run(args []string, env env.ShellEnvironment) error {
    final := "\nGrove Help:\n-----\n"

    legend := "Legend:\n- [\"example\"] -> string value (either normal or raw)\n- [value1/value2] -> either value1 or value (type-wise)\n- [example]... -> infinite amount of arguments\n- other syntax is presumed to be self-explanatory\n"
    final += legend

    final += "\nCommands:\n-----\n"
    for _, command := range Commands {
        final += command.Usage() + "\n - " + command.Description() + "\n"
    }

    fmt.Println(final)

    return nil
}
