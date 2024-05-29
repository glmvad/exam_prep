package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2] 

	var i,j int
	res := ""

	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			res += string(str1[i])
			i++
		}
		j++ 
	}

	if i == len(str1) {
		fmt.Println(res)
	}
	
	


	
	
}
