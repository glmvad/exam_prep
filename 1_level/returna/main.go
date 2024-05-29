package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	flag := false 
	for _, v := range str {
		if v == 'a' {
			flag = true 
			
			break
		} else {
			flag = false 
			
		}
	}
	if flag {
		fmt.Print("Contains the letter\n")
	} else {
		fmt.Print("!(Contains the letter)\n")
	}

}