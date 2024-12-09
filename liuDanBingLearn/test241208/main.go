package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	arr := []int{1, 2, 4, 3}
	slices.SortFunc(arr, func(a, b int) int {
		return b - a
	})
	fmt.Println(arr)

	s := "1345"
	num, _ := strconv.Atoi(s)
	fmt.Println(num)
}
