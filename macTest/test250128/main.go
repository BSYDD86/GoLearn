package main

import (
	"fmt"
	"strconv"
)

func main() {
	var res string
	a := "a"
	res = min(res, a)
	fmt.Println(res)
}
func maximumSwap(num int) int {
	val := strconv.Itoa(num)
	arr := []byte(val)
	for i := 0; i < len(arr)-1; i++ {
		next := i
		for j := 1; j < len(arr); j++ {
			if arr[j] > arr[next] {
				next = j
			}
		}
		if next != i {
			arr[i], arr[next] = arr[next], arr[i]
			break
		}
	}
	str := string(arr)
	res, _ := strconv.Atoi(str)
	return res
}

func firstDayBeenInAllRooms(nextVisit []int) int {
	res := 0
	cnt := len(nextVisit)
	m := make([]int, len(nextVisit))
	n := len(nextVisit)
	nextIdx := 0
	for cnt > 0 {
		if m[nextIdx] == 0 {
			cnt--
			m[nextIdx]++
		}
		res++
		if res%2 == 0 {
			nextIdx = (nextIdx + 1) % n
		} else {
			nextIdx = nextVisit[nextIdx]
		}

	}
	return res
}

func dfs(grid [][]byte, row, col int) {
	path := [][]int{{-1, 0}, {0, -1}, {1, 0}, {1, 0}}
	rows := len(grid)
	cols := len(grid[0])
	var isValid func(rowIdx int, colIdx int, rows int, cols int) bool
	isValid = func(rowIdx int, colIdx int, rows int, cols int) bool {
		if rowIdx < 0 || rowIdx >= rows {
			return false
		}
		if colIdx < 0 || colIdx >= cols {
			return false
		}
		return true
	}
	grid[row][col] = '0'
	for i := 0; i < 4; i++ {
		nextRowIdx := row + path[i][0]
		nextColIdx := col + path[i][1]
		if !isValid(nextRowIdx, nextColIdx, rows, cols) || grid[nextRowIdx][nextColIdx] != '1' {
			continue
		} else {
			dfs(grid, nextRowIdx, nextColIdx)
		}
	}
}
