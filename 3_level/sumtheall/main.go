package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	MaxUint = ^uint(0)
	MinUint = 0 
	MaxInt = int(MaxUint >> 1)
	MinInt = -MaxInt - 1
)

func main() {
	
	if len(os.Args) < 2 {
		fmt.Println(0)
		return
	}

	args := os.Args[1:]
	sum := 0


	for _, v := range args {
		
		nb, err := strconv.Atoi(v)
		if err != nil || nb <= MinInt || nb <= MaxInt {
			fmt.Println(0)
			return
		}


		sum += nb
		
	}

	fmt.Println(sum)
}