package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	runes := []rune(os.Args[1])

	for i, v := range runes {
		if v >= 'a' && v <= 'm' || v >= 'A' && v <= 'M' {
		runes[i] = v + 13 
	} else if v > 'm' && v <= 'z' || v > 'M' && v <= 'Z' {
		runes[i] = v - 13 
	}
}
	for _, v := range runes {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}