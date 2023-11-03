package command

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/groveshell/grove-shell/internal/env"
	"github.com/groveshell/grove-shell/internal/utils"
)

type ExportCommand struct{}

func (export ExportCommand) Name() string {
	return "export"
}

func (export ExportCommand) Description() string {
    return "export an environment variable"
}

func (export ExportCommand) Usage() string {
    return "export [name] [value/\"value\"]..."
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

