package cmd

import (
	"fmt"
	"os"
	"strconv"

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

func (cd CdCommand) Run(args []string) error {
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

func (echo EchoCommand) Run(args []string) error {
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

func (exit ExitCommand) Run(args []string) error {
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

func (pwd PWDCommand) Run(args []string) error {
    wd, err := os.Getwd()
    if err != nil {
        return err
    }
    
    fmt.Println(wd)

    return nil
}
