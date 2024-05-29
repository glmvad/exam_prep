package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return 
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	res := ""

	for _, v := range s1 {
		if !Contain(res, v) {
			res += string(v)
		}
	}

	for _, v := range s2 {
		if !Contain(res, v) {
			res += string(v)
		}
	}
	
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Contain(res string, r rune) bool {
	for _, v := range res {
		if v == r {
			return true 
		}
	}
	return false
}