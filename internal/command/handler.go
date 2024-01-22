package command

import "github.com/mattishere/grove-shell/internal/env"

type CommandHandler struct {
	cmds map[string]Command
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		cmds: make(map[string]Command),
	}
}

func (ch CommandHandler) RegisterNew(cmd Command) {
	ch.cmds[cmd.Name()] = cmd
}

func (ch CommandHandler) RunCmd(name string, args []string, env env.ShellEnvironment) (doesExist bool, err error) {
	cmd, exists := ch.cmds[name]
	if !exists {
		return false, nil
	}

	err = cmd.Run(args, env)
	if err != nil {
		return true, err
	}

	return true, nil
}
