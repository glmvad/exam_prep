package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2  {
		return
	}

	s := os.Args[1]
	res := ""
	
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res += string(s[i])
		} else if i != 0 && s[i] == ' ' && s[i-1] >= 'a' && s[i-1] <= 'z'  {
			
			n := 3
			for n > 0 {
				res += string(s[i])
				n--
			}
		
		}  

		
		}

		if res[len(res)-1] == ' ' {
			res = res[:len(res)-3]
		}
		

fmt.Println(res)
}