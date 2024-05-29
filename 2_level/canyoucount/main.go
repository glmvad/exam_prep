package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("0")
	}
	args := os.Args[1:]
	res := 0
	for _, v := range args{
		res += len(v)
	}
	fmt.Println(res)
}