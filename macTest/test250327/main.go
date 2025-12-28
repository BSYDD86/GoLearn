package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "3[a2[c]]"
	fmt.Println(decodeString(str))
}
func decodeString(s string) string {
	cnt := 0
	res := ""
	st := []string{}
	stCnt := []int{}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			temp, _ := strconv.Atoi(string(c))
			cnt = cnt*10 + temp
		} else if c >= 'a' && c <= 'z' {
			res += string(c)
		} else if c == '[' {
			st = append(st, res)
			res = ""
			stCnt = append(stCnt, cnt)
			cnt = 0
		} else {
			tempCnt := stCnt[len(stCnt)-1]
			stCnt = stCnt[:len(stCnt)-1]
			res = strings.Repeat(res, tempCnt)
			res = st[len(st)-1] + res
			st = st[:len(st)-1]
		}
	}
	return res
}

//	func main() {
//		arr := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
//		println(largest1BorderedSquare(arr))
//	}
func largest1BorderedSquare(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	//[0,i-1]
	right := make([][]int, rows+1)
	for i := range right {
		right[i] = make([]int, cols+1)
	}
	//[0,j-1]
	down := make([][]int, rows+1)
	for i := range down {
		down[i] = make([]int, cols+1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := grid[i][j]
			down[i+1][j] = down[i][j] + x
			right[i][j+1] = right[i][j] + x
		}
	}
	d := min(rows, cols)
	for d > 0 {
		for i := 0; i+d <= rows; i++ {
			for j := 0; j+d <= cols; j++ {
				if right[i][j+d]-right[i][j] == d &&
					right[i+d-1][j+d]-right[i+d-1][j] == d &&
					down[i+d][j]-down[i][j] == d &&
					down[i+d][j+d-1]-down[i][j+d-1] == d {
					return d * d
				}
			}
		}
		d--
	}
	return 0
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	path := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	cnt := 0
	time := 0
	stOrange := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				cnt++
			} else if grid[i][j] == 2 {
				stOrange = append(stOrange, []int{i, j})
			}
		}
	}
	for cnt > 0 && len(stOrange) > 0 {
		tempSize := len(stOrange)
		for i := 0; i < tempSize; i++ {
			tempRow := stOrange[i][0]
			tempCol := stOrange[i][1]
			for j := 0; j < len(path); j++ {
				nextRow := tempRow + path[j][0]
				nextCol := tempCol + path[j][1]
				if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
					continue
				} else {
					if grid[nextRow][nextCol] == 1 {
						grid[nextRow][nextCol] = 2
						cnt--
						stOrange = append(stOrange, []int{nextRow, nextCol})
					}
				}
			}
		}
		time++
		stOrange = stOrange[tempSize:]
	}
	if cnt != 0 {
		return -1
	} else {
		return time
	}
}
