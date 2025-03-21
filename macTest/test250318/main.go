package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := make([]int, 2, 4)
	t := arr[1:]
	t = append(arr, 1)
	t[0] = 12
	t = append(t, 2)
	t[0] = 13
	t = append(t, 2)
	t[0] = 14
	t = append(t, 2)
	t[0] = 11
	t = append(t, 2)
	fmt.Println(cap(t))
	fmt.Println(len(t))
}

func decodeString(s string) string {
	res := ""
	cnt := 0

	stStr := []string{}
	stCnt := []int{}
	for _, c := range s {
		if '0' <= c && c <= '9' {
			val, _ := strconv.Atoi(string(c))
			cnt *= 10
			cnt += val
		} else if c == '[' {
			stStr = append(stStr, res)
			res = ""
			stCnt = append(stCnt, cnt)
			cnt = 0
		} else if c == ']' {
			tempCnt := stCnt[len(stCnt)-1]
			stCnt = stCnt[:len(stCnt)-1]
			res = strings.Repeat(res, tempCnt)
			res = stStr[len(stStr)-1] + res
			stStr = stStr[:len(stStr)-1]

		} else {
			res = res + string(c)
		}
	}
	return res
}

var (
	path [][]byte
	res  [][]string
)

func solveNQueens(n int) [][]string {
	res = [][]string{}
	path = make([][]byte, n)
	for i := range path {
		path[i] = make([]byte, n)
		for j := range path {
			path[i][j] = '.'
		}
	}
	dfs(n, 0, 0)
	return res
}

func dfs(n int, row int, cnt int) {
	if row == n && cnt == n {
		t := make([]string, 0, n)
		for i := range path {
			s := string(path[i])
			if s != "" {
				t = append(t, s)
			}
		}
		res = append(res, t)
		return
	}
	for j := 0; j < n; j++ {
		if check(n, row, j) {
			path[row][j] = 'Q'
			dfs(n, row+1, cnt+1)
			path[row][j] = '.'
		}
	}
}

// is Legal
func check(n int, row int, col int) bool {
	for i := row - 1; i >= 0; i-- {
		if path[i][col] == 'Q' {
			return false
		}
	}
	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		if path[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	i = row - 1
	j = col + 1
	for i >= 0 && j < n {
		if path[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	return true
}
