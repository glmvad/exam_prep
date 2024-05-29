package main

import (
	"fmt"
)

func main() {
	fmt.Println(ReverseSecondHalf("This is the 1st half This is the 2nd halmf"))
	fmt.Println(ReverseSecondHalf(""))
	fmt.Println(ReverseSecondHalf("Hello World"))
}

func ReverseSecondHalf(str string) string {
	if len(str) == 0 {
		return "Invalid Output"
	}
	
	for i := 1; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str = str[i:]
			break 
		}
	}
	res := ""
	
	
	if len(str)%2 != 0 {
		for i := len(str)-1 ; i >=0; i-- {
			res += string(str[i]) 
		}
	} else {
		return str
	}
	return res
}
