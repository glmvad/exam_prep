package main 

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	res := ""
	flag := true 

	for i, v := range s {
		if i!=0 && (s[i] >= 'A' && s[i] <= 'Z') && (s[i-1] >= 'a' && s[i-1] <= 'z')   {
			 
			res += "_" + string(v) 
		} else if (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z')  {
			flag = false 
			break 
		} else {
			res += string(v)
		}
	}
	if flag {
		return res 
	} else {
		return s
	}
}