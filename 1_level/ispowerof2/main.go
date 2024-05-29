package main

import (
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func main() {
if len(os.Args) != 2 {
	return
}

n1, _ := strconv.Atoi(os.Args[1])
n2 := 2 

for n1 > n2 {
	n2 *= 2
}
print := []rune{}
if n2 > n1  {
	print = []rune{'f', 'a', 'l', 's', 'e'}
} else {
	print = []rune{'t', 'r', 'u', 'e'}
}

for _, v := range print {
	z01.PrintRune(v)
} 
z01.PrintRune('\n')

}
