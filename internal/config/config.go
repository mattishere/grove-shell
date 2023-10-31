package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Aliases       map[string]string `toml:"aliases"`
	Customization Customization     `toml:"customization"`
}

type Customization struct {
	Prompt string `toml:"prompt"`
}

var (
	defaultConfig = Config{
		Customization: Customization{Prompt: "$ "},
	}
)

func GetConfig() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return defaultConfig
	}

	path := filepath.Join(homeDir, ".groverc")

	_, err = os.Stat(path)
	if err != nil {
		return defaultConfig
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return defaultConfig
	}

	var config Config
	if err := toml.Unmarshal(content, &config); err != nil {
		return defaultConfig
	}
	return config
}
