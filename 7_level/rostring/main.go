package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	word := ""
	words := []string{}

	for i := 0; i < len(arg); i++ {
		if arg[i] != ' ' {
			word += string(arg[i])
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
		if i == len(arg)-1 && word != "" {
			words = append(words, word)
		}
	}

	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			z01.PrintRune(rune(words[i][j]))
		}
		z01.PrintRune(' ')
	}

	for _, v := range words[0] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')

}
