package main

import (
	"fmt"
	"sort"
)

// sortTest
func main() {
	res := [][]int{{1, 3}, {3, 1}, {2, 2}}

	sort.Slice(res, func(i, j int) bool {
		return res[i][1] < res[j][1]
	})

	fmt.Println(res)

}
