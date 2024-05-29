package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	nb, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	
	res := ""
	for i:=2; i <= nb; i++ {
		for nb > 0 && nb%i == 0  {
				nb /= i
				res += strconv.Itoa(i) + "*"
			
		}
	}
	fmt.Println(res[:len(res)-1])
	
}

