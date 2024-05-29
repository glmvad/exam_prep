package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	runes1 := []rune(os.Args[1])
	r2 := os.Args[2][0]
	r3 := os.Args[3][0]

	for i := 0; i < len(runes1); i++ {
		if runes1[i] == rune(r2) {
			runes1[i] = rune(r3) 
		}
	}
	for _, v := range runes1 {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}