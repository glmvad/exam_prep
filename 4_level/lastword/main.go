package main

import (
	"fmt"
	"os"
)

// func main() {
// 	s := os.Args[1]
// 	res := ""

// 	for i:= len(s)-1 ; i >= 0; i-- {
// 		if s[i] != ' ' {
// 			res = string(s[i]) + res 
// 		} else if res != "" {
// 			break 
// 		}
// 	}
// 	fmt.Println(res)
// }


































func main() {
	if len(os.Args) != 2 {
		return 
	}
	arg := os.Args[1]
	word := ""
	for i := len(arg)-1; i >= 0; i-- {
		if arg[i] != ' ' {
			word = string(arg[i]) + word
		} else if word != "" {
			break
		}
	}

	fmt.Println(word)

}
