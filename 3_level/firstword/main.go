package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return 
	}

	arg := os.Args[1]

	w := ""

	for i := 0; i < len(arg)-1; i++ {
		if arg[i] != ' ' {
			w += string(arg[i])
		} else {
			break
		}
	}

	fmt.Println(w)
}