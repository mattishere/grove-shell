package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/groveshell/grove-shell/internal/run"
)

func StartShell() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("# ")
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal("Failed to read input.")
        }

        hasQuit, err := run.RunCommand(input)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        if hasQuit {
            break
        }
    }
}
