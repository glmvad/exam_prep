package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2{
		return
	}
	
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Stdout.WriteString("ERROR\n")
		return
	}

	a := "0123456789abcedf"
	res := ""
	temp := 0

	for n > 0 {
		temp = n%16
		res = string(a[temp]) + res 
		n /= 16
	}

	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}