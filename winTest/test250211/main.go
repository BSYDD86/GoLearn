package main

import "fmt"

func main() {
	m := map[int]int{}
	m[1] = 1
	fmt.Println(m[1])
	delete(m, 1)
	fmt.Println(m[1])
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < (n+1)/2; j++ {
			// i,j
			// n-1-j,i

			//i,j
			//n-1-j,i
			//n-1-i,n-1-j
			//j,n-1-i
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}
