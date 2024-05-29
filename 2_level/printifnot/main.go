package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}

func PrintIfNot(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}