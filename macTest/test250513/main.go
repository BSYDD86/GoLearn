package main

import "fmt"

func main() {
	i := int(1)
	res := ""
	res += string(byte(i))
	fmt.Println(res)
	fmt.Println(byte(i))
}
