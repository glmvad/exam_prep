package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	a, _ :=  strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])

	temp := 0 
	gcd := 0 

	if a > b {
		temp = a
	} else {
		temp = b
	}



	for i := 1; i <= temp; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	} 
	fmt.Println(gcd)
}