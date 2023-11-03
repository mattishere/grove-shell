package lex

func Lex(input string) []string {
	var tokens []string
	current := ""
	var isString bool

	for _, char := range input {
		switch char {
		case ' ', '\r', '\t':
			if current != "" {
				if !isString {
					tokens = append(tokens, current)
					current = ""
				} else {
					current += string(char)
				}
			}
		case '"', '\'':
			isString = !isString
			current += string(char)
		case '#':
			if current == "" && !isString {
				return tokens
			} else {
				current += string(char)
			}
		default:
			current += string(char)
		}
	}

	if current != "" {
		tokens = append(tokens, current)
	}

	return tokens
}
