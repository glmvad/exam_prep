package main

import (
    "fmt"
    "os"
)

func main() {

	if len(os.Args) == 0 {
		return 
	}
	args := os.Args[1:]
    for _, v := range args {
        if isBalanced(v) {
            fmt.Println("ok")
        } else {
            fmt.Println("error")
        }
    }
}

func isBalanced(arg string) bool {
    slice := []rune{}
    brackets := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
	
    for _, v := range arg {
        if v == '{' || v == '(' || v == '[' {
            slice = append(slice, v)
        } else if v == '}' || v == ')' || v == ']' {
            if len(slice) == 0 || slice[len(slice)-1] != brackets[v] {
                return false
            }
			
            slice = slice[:len(slice)-1]
        }
    }
   
	return len(slice)==0

}



