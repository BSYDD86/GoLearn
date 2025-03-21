package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3}
	res := append(arr[:1], arr[2:]...)
	fmt.Println(res)
	fmt.Println(sort.SearchInts(arr, 1))
	fmt.Println(sort.SearchInts(arr, 4))

}
