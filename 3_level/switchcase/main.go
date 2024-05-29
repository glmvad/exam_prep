package main

import (
	
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return 
	}

	s := os.Args[1]

	res := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			
			res += string(s[i]+32)

	} else {
		
		res += string(s[i]-32)
	}
}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}