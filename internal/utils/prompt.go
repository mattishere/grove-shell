package utils

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func GeneratePrompt(prompt string) (string, error) {
	placeholders := map[string](func() (string, error)){
		"username": func() (string, error) {
			currUser, err := user.Current()
			return currUser.Username, err
		},
		"hostname": func() (string, error) {
			host, err := os.Hostname()
			return host, err
		},
		"path": func() (string, error) {
			wd, err := os.Getwd()
			return wd, err
		},
		"curr_dir": func() (string, error) {
			wd, err := os.Getwd()
			return filepath.Base(wd), err
		},
	}

	styles := map[string]int{
		"reset":     0,
		"bold":      1,
		"underline": 4,
		"black":     30,
		"red":       31,
		"green":     32,
		"yellow":    33,
		"blue":      34,
		"magenta":   35,
		"cyan":      36,
		"white":     37,
	}

	final := prompt

	for placeholder, getValue := range placeholders {
		value, err := getValue()
		if err != nil {
			return "", err
		}

		final = strings.ReplaceAll(final, "{"+placeholder+"}", value)
	}

	for style, value := range styles {
		fullStyle := "\033[" + strconv.Itoa(value) + "m"

		final = strings.ReplaceAll(final, "{"+style+"}", fullStyle)
	}

	return final, nil
}

func DefaultPrompt() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	host, err := os.Hostname()
	if err != nil {
		return "", err
	}

	currUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return "[" + currUser.Username + "@" + host + ":" + wd + "]$ ", nil
}
