package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/groveshell/grove-shell/internal/env"
	"github.com/groveshell/grove-shell/internal/utils"
)

type Command interface {
	Name() string
	Run([]string, env.ShellEnvironment) error
}

type CdCommand struct{}

func (cd CdCommand) Name() string {
	return "cd"
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
		fmt.Println(path)
	}

	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("path error: %s not found", path)
	}

	return nil
}

type EchoCommand struct{}

func (echo EchoCommand) Name() string {
	return "echo"
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

type ExitCommand struct{}

func (exit ExitCommand) Name() string {
	return "exit"
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

type PWDCommand struct{}

func (pwd PWDCommand) Name() string {
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

type ExportCommand struct{}

func (export ExportCommand) Name() string {
	return "export"
}

func (export ExportCommand) Run(args []string, env env.ShellEnvironment) error {
	if len(args) >= 2 {
		name := args[0]

		if !unicode.IsLetter(rune(name[0])) && rune(name[0]) != '_' {
			return fmt.Errorf("invalid variable name")
		}

		for _, char := range name[1:] {
			if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
				return fmt.Errorf("invalid variable name")
			}
		}

		values := args[1:]

		var finalValues []string
		for _, value := range values {
			if utils.IsString(value) || utils.IsRawString(value) {
				finalValues = append(finalValues, value[1:len(value)-1])
			} else {
				finalValues = append(finalValues, value)
			}
		}

		final := strings.Join(finalValues, " ")
		err := os.Setenv(name, final)
		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("not enough arguments")
}

type AliasCommand struct{}

func (alias AliasCommand) Name() string {
	return "alias"
}

func (alias AliasCommand) Run(args []string, env env.ShellEnvironment) error {
	if len(args) == 2 {
		aliasKey := args[0]
        if utils.IsString(aliasKey) || utils.IsRawString(aliasKey) {
            return fmt.Errorf("invalid alias name (cannot be string)")
        }

        command := args[1]
        if utils.IsString(command) || utils.IsRawString(command) {
            command = command[1:len(command)-1]
        }
        env.Aliases[aliasKey] = command
    } else if len(args) == 0 {
        for name, alias := range env.Aliases {
            fmt.Println(name + " -> " + alias)
        }
    }

    return nil
}
