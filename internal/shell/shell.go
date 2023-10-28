package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/groveshell/grove-shell/internal/cmd"
	"github.com/groveshell/grove-shell/internal/run"
)

func StartShell() {
	reader := bufio.NewReader(os.Stdin)

	cmdHandler := cmd.NewCommandHandler()
	cmdHandler.RegisterNew(cmd.CdCommand{})
	cmdHandler.RegisterNew(cmd.EchoCommand{})
	cmdHandler.RegisterNew(cmd.ExitCommand{})

	for {
		fmt.Print("-> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed to read input.")
		}

        input = strings.TrimSpace(input)

        if input == "" {
            continue
        }

		err = run.RunCommand(cmdHandler, input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
