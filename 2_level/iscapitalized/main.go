package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4his 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {

	flag := false
	for i := 0; i < len(s); i++ {
		if i==0 && (s[i] >= 'A' && s[i] <= 'Z') || i==0 && s[i] < 65 {
			flag = true
		} else if i != len(s)-1 && s[i] == ' ' && (s[i+1] >= 'a' && s[i+1] <= 'z'){
			flag = false 
			break
		} else {
			continue
		}
	}
	if flag {
		return true 
	} else {
		return false
	}
}