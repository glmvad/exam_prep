package main

import (
	"os"
)

func main() {
	arg := os.Args[1]

	if len(os.Args) != 2 {
		return
	}

	at := -1
	res := ""

	for i, v := range arg {
		if IsVowel(v) {
			at = i
			break
		} else {
			res += string(v)
		}
	}
	
	if at > -1 {
		os.Stdout.WriteString(arg[at:]+res+"ay\n")
	} else {
		os.Stdout.WriteString("No vowels\n")
}
}




func IsVowel(r rune) bool {
	vowel := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range vowel {
		if v == r {
			return true 
		}
	}
	return false
}