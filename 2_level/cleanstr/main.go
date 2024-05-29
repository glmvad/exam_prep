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
	if arg == "" {
		return
	}
	res := ""
	for i := 0; i < len(arg); i++ {
		if (i == 0 || i == len(arg)-1) && arg[i] == ' ' {
			continue
		} else if i != len(arg)-1 && arg[i] == ' ' && arg[i+1] == ' ' {
			continue
		} else {
			res += string(arg[i])
		}
	}
	fmt.Println(res)
}