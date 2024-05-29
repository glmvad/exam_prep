package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	res := ""
	temp := 0
	for _, v := range dec {
		temp = (len(dec) + int(v))%127
		res += string(temp) 
	}
	return res
}
