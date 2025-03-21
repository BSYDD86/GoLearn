package main

import "fmt"

func main() {
	arr := []int{1, 2}
	res := []int{1, 3}
	res = arr
	fmt.Println(res)
}
