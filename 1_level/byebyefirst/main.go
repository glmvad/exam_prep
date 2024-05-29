package main

import (
	"fmt"
)

func main() {
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
	fmt.Println(ByeByeFirst([]string{"abcd", "efg", "ff"}))
}

func ByeByeFirst(strings []string) []string {

	if len(strings) < 2 {
		return []string{}
	}
	slice := []string{}

	for i := 1; i < len(strings); i++ {
		slice = append(slice, strings[i])
	}
	return slice
}