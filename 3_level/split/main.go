package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	arr := []string{}
	w := ""

	for i := 0; i< len(s); i++ {
		if s[i] == sep[0] && s[i+1] == sep[1] && i != len(s)-1  {
			arr = append(arr, w)
			w = ""
			i ++
		} else {
			w += string(s[i]) 
		}
	}
	arr = append(arr, w)
	return arr

}
