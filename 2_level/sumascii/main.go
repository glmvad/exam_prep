package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
	}

	a := os.Args[1]
	sum := 0

	for _, v := range a {
		sum += int(v)
	}
	sum %= 256
	fmt.Println(sum)
}