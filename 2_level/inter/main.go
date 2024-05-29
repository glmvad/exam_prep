package main

import (
	
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	res := ""

	m := make(map[string]bool)

	for i := 0; i < len(str2); i++ {
		m[string(str2[i])] = true
	}

	for i := 0; i < len(str1); i++ {
		if m[string(str1[i])] {
			res += string(str1[i])
			delete(m, string(str1[i]))
		}
	}

	for _, v := range res {
		z01.PrintRune(v)
	}
	
	z01.PrintRune('\n')
}