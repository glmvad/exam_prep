package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println()
	}
	res := ""

	for _, v := range args {
		
		for i := 0; i < len(v); i++ {
			if i%2 == 0 {
				res += "2"
			} else {
				res += string(v[i])
			}
		}	
		res += " "
	}
	fmt.Println(res[:len(res)-1])
}
