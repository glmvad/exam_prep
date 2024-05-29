package main

import (
	"fmt"
)

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	neg := ""
	temp := 0
	res := ""
	if n < 0 {
		n = -n 
		neg = "-"
	}

	if n == 0 {
		return "0"
	}

	for n > 0 {
		temp = n%10 
		res = string(temp+48) + res 
		n /= 10 
	}

	return neg+res
}

