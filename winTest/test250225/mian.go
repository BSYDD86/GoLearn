package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(arr)
	fmt.Println(arr)
}

func rotate(matrix [][]int) {
	//i,j
	//n-1-j,i
	//n-1-i,n-1-j
	//j,n-1-i
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2+1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}

	}
}
