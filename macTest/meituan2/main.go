package main

import "fmt"

//6
//0 0
//0 1
//0 2
//1 0
//2 0
//3 0

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
		if x < 0 || y < 0 {

		}
		arr[x][y] = true
		resTemp[i][0] = x
		resTemp[i][1] = y
	}
	for i := 0; i < n; i++ {
		x := resTemp[i][0]
		y := resTemp[i][1]
		res := 0

		//up
		temp := 0
		for i := x - 1; i >= 0 && temp <= 1; i-- {
			if arr[i][y] {
				temp++
			}
		}
		if temp > 1 {
			res++
		}
		//down
		temp = 0
		for i := x + 1; i < n && temp <= 1; i++ {
			if arr[i][y] {
				temp++
			}
		}
		if temp > 1 {
			res++
		}
		//left
		temp = 0
		for j := y - 1; j >= 0 && temp <= 1; j-- {
			if arr[x][j] {
				temp++
			}
		}
		if temp > 1 {
			res++
		}
		//right
		temp = 0
		for j := y + 1; j < n && temp <= 1; j++ {
			if arr[x][j] {
				temp++
			}
		}
		if temp > 1 {
			res++
		}
		fmt.Println(res)
	}

}
