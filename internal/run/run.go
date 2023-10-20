package run

import (
	"os"
	"os/exec"
	"strings"
)

func RunCommand(input string) (bool, error) {
    split := strings.Split(strings.TrimSpace(input), " ")
    command := split[0]
    args := split[1:]

    if command == "exit" {
        return true, nil
    }

    cmd := exec.Command(command, args...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        return false, err
    }

    return false, nil
}
