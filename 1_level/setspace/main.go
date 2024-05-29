package main

import (
	"fmt"
)

func main() {
	fmt.Println(SetSpace("HelloWorld"))
	fmt.Println(SetSpace("HelloWorld12"))
	fmt.Println(SetSpace("Hello World"))
	fmt.Println(SetSpace(""))
	fmt.Println(SetSpace("LoremIpsumWord"))
}

func SetSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := ""

	for i := 0; i < len(s); i++ {
		if i != 0 && s[i] >= 'A' && s[i] <= 'Z' {
			res += string(' ') + string(s[i])
		} else if s[i] == ' ' || s[i] < 65 || s[i] > 122 {
			return "Error"
		} else {
			res += string(s[i])
		}
	}
	return res
}
