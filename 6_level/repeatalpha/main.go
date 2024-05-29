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

	for j := 0; j < len(arg); j++ {
		if arg[j] >= 'A' && arg[j] <= 'Z' {
			for i := 0; i < int(arg[j]-64); i++ {
				z01.PrintRune(rune(arg[j]))
			}
		} else if arg[j] >= 'a' && arg[j] <= 'z' {
			for i := 0; i < int(arg[j]-96); i++ {
				z01.PrintRune(rune(arg[j]))
			}
		} else {
			z01.PrintRune(rune(arg[j]))
		}
	}

	z01.PrintRune('\n')
}
