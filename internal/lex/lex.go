package lex

func Lex(input string) []string {
	var tokens []string
	current := ""
	var isString, isRawString, isEscaped bool

	for _, char := range input {
		switch char {
		case ' ', '\r', '\t':
			if current != "" {
				if !isString && !isRawString {
					tokens = append(tokens, current)
					current = ""
				} else {
					current += string(char)
				}
			}
		case '"':
			if isRawString || isEscaped {
				current += string(char)
				isEscaped = false
			} else {
				isString = !isString
				current += string(char)
			}
		case '\'':
			if isString {
				current += string(char)
			} else {
				isRawString = !isRawString
				current += string(char)
			}
		case '\\':
			if isString {
				isEscaped = true
			} else {
				current += string(char)
			}
		case '#':
			if current == "" && !isString && !isRawString {
				return tokens
			} else {
				current += string(char)
			}
		default:
			current += string(char)
			isEscaped = false
		}
	}

	if current != "" {
		tokens = append(tokens, current)
	}

	return tokens
}
