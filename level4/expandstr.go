package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	words := make([]string, 0)
	currentWord := ""

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(input[i])
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	if len(words) == 0 {
		return
	}

	for i, word := range words {
		for _, char := range word {
			z01.PrintRune(char)
		}
		if i < len(words)-1 {
			for j := 0; j < 3; j++ {
				z01.PrintRune(' ')
			}
		}
	}

	z01.PrintRune('\n')
}