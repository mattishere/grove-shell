package expand

import (
	"os"
	"strings"
)

type Expander interface {
	Expand(input string) string
}

type EnvironmentExpander struct{}

func (ee *EnvironmentExpander) Expand(input string) string {
	return os.ExpandEnv(input)
}

type HomeDirExpander struct{}

func (hde *HomeDirExpander) Expand(input string) string {
    out := input
	for i, char := range input {
		if char == '~' && (i == 0 || input[i-1] == ' ') {
			if (i+1 < len(input) && (input[i+1] == ' ' || input[i+1] == '/')) || i+1 == len(input) {
                home, err := os.UserHomeDir()
                if err == nil {
                    out = strings.Replace(out, "~", home, 1)
                }
			}
		}
	}

    return out
}
