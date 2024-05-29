package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	var i, j int
	res := ""

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			res += string(s1[i])
			i++
		}
		j++
	}

	if len(s1) == len(res) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
