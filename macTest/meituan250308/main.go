package main

import "fmt"

func main() {

	//test2
	var n int
	fmt.Scan(&n)
	arr := make([][]bool, n)

	resTemp := make([][]int, n)
	for i := range resTemp {
		resTemp[i] = make([]int, 2)
	}
	for i := range arr {
		arr[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		arr[x][y] = true
		resTemp[i][0] = x
		resTemp[i][1] = y
	}

	dpl := make([][]int, n)
	for i := range dpl {
		dpl[i] = make([]int, n)
	}
	dpr := make([][]int, n)
	for i := range dpr {
		dpr[i] = make([]int, n)
	}

	dpu := make([][]int, n)
	for i := range dpu {
		dpu[i] = make([]int, n)
	}
	dpd := make([][]int, n)
	for i := range dpd {
		dpd[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		temp := 0
		for j := 0; j < n; j++ {
			if arr[i][j] {
				temp++
				dpl[i][j] = temp
			}
		}
	}
	for i := 0; i < n; i++ {
		temp := 0
		for j := n - 1; j >= 0; j-- {
			if arr[i][j] {
				temp++
				dpr[i][j] = temp
			}
		}
	}

	for j := 0; j < n; j++ {
		temp := 0
		for i := 0; i < n; i++ {
			if arr[i][j] {
				temp++
				dpu[i][j] = temp
			}
		}
	}
	for j := 0; j < n; j++ {
		temp := 0
		for i := n - 1; i >= 0; i-- {
			if arr[i][j] {
				temp++
				dpd[i][j] = temp
			}
		}
	}

	for i := 0; i < n; i++ {
		x := resTemp[i][0]
		y := resTemp[i][1]
		res := 0
		if dpl[x][y] > 2 {
			res += 1
		}
		if dpr[x][y] > 2 {
			res += 1
		}
		if dpu[x][y] > 2 {
			res += 1
		}
		if dpd[x][y] > 2 {
			res += 1
		}
		fmt.Println(res)
	}
}
