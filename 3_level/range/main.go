package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	for i := a; i <= b; i++ {
		
		fmt.Print(i, " ")
	}
	fmt.Println()
}
