package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
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
    cmdHandler.RegisterNew(cmd.PWDCommand{})


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



		err = run.RunCommand(cmdHandler, input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Prompt() string {
    wd, err  := os.Getwd()
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
