package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mattishere/grove-shell/internal/command"
	"github.com/mattishere/grove-shell/internal/env"
	"github.com/mattishere/grove-shell/internal/files"
	"github.com/mattishere/grove-shell/internal/run"
	"github.com/mattishere/grove-shell/internal/utils"
)

func StartShell() {
	cmdHandler := command.NewCommandHandler()
	for _, cmd := range command.Commands {
		cmdHandler.RegisterNew(cmd)
	}

	env := env.ShellEnvironment{
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
	rc := files.ReadRCFile()
	if len(rc) > 0 {
		for _, line := range rc[0 : len(rc)-1] {
			run.RunCommand(cmdHandler, line, env)
		}
	}

	for {
		prompt, err := utils.DefaultPrompt()
		if err != nil {
			prompt = "$ "
		}

		if value, exists := os.LookupEnv("PROMPT"); exists {
			generatePrompt, err := utils.GeneratePrompt(value)
			if err == nil {
				prompt = generatePrompt
			}
		}

		fmt.Print(prompt)
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
