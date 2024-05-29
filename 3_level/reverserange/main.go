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

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("strconv.Atoi: parsing %s: invalid syntax\n", err)
		return
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("strconv.Atoi: parsing %s: invalid syntax\n", err)
		return
	}

	for i := b; i >= a; i-- {
		fmt.Print(i)
		if i == a {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}	
}