package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/groveshell/grove-shell/internal/cmd"
	"github.com/groveshell/grove-shell/internal/env"
	"github.com/groveshell/grove-shell/internal/files"
	"github.com/groveshell/grove-shell/internal/run"
)

func StartShell() {
	cmdHandler := cmd.NewCommandHandler()
	cmdHandler.RegisterNew(cmd.CdCommand{})
	cmdHandler.RegisterNew(cmd.EchoCommand{})
	cmdHandler.RegisterNew(cmd.ExitCommand{})
	cmdHandler.RegisterNew(cmd.PWDCommand{})
	cmdHandler.RegisterNew(cmd.ExportCommand{})
    cmdHandler.RegisterNew(cmd.AliasCommand{})

    env := env.ShellEnvironment{
        Variables: make(map[string]string),
        Aliases: make(map[string]string),
    }

	if len(os.Args) > 1 {
		lines, err := files.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		for _, line := range lines[0 : len(lines)-1] {
			err = run.RunCommand(cmdHandler, line, env)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(Prompt())
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed to read input.")
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		err = run.RunCommand(cmdHandler, input, env)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Prompt() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	currUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	return "[" + currUser.Username + "@" + host + ":" + wd + "]$ "
}
