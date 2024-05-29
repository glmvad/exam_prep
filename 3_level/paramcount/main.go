package main

import (

	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		os.Stdout.WriteString("0\n")
	} 
	
	count := Itoa(len(args))
	
	for _, v := range count {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	temp := 0
	res := ""

	for n > 0 {
		temp = n%10 
		res = string(temp+48) + res 
		n /= 10 
	}
	return res 
}
