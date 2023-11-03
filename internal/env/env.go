package env

type ShellEnvironment struct {
	Variables map[string]string
	History   []string
	Aliases   map[string]string
}
