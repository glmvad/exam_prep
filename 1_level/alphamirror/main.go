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

	for i:=0; i < len(runes); i++ {
	if runes[i] >= 'a' && runes[i] <= 'z' {
		runes[i] = 'a' + 'z' - runes[i] 
	}
	if runes[i] >= 'A' && runes[i] <= 'Z' {
		runes[i] = 'A' + 'Z' - runes[i] 
	}
}
for _, v := range runes {
	z01.PrintRune(v)
}
z01.PrintRune('\n')
}