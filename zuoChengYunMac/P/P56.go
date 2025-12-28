package main

import (
	"fmt"
)

func main() {
	var n int
	var m int
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		return
	}
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, 3)
	}
	for i := range m {
		_, err := fmt.Scan(&arr[i][0], arr[i][1], arr[i][2])
		if err != nil {
			return
		}
	}
}

type DisJointSet struct {
	father []int
	size   []int
}

func (d *DisJointSet) find(a, b int) bool {
	return d.father[a] == d.father[b]
}
