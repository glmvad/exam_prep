package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2  {
		return
	}

	s := os.Args[1]
	w := ""
	arr := []string{}

	for i := len(s)-1; i >= 0; i-- {
		if s[i] != ' ' {
			w = string(s[i]) + w
		} else if  s[i] == ' ' {
			arr = append(arr, w)
			w = ""
		} 
	}
	
	arr = append(arr, w)

	res := ""
	for i := 0; i < len(arr); i++ {
		res += arr[i] + " "
	} 

	res = res[:len(res)-1]
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
