package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	words := make([]rune, 0, len(s))
	for _, c := range s {
		if c == ' ' {
			if len(words) > 0 && !isUpper(words[0]) && isLetter(words[0]) {
				return false
			}
			words = make([]rune, 0, len(s))
		} else {
			words = append(words, c)
		}
	}

	if len(words) > 0 && !isUpper(words[0]) && isLetter(words[0]) {
		return false
	}

	return true
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}