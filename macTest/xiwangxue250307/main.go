package main

func main() {

}

var (
	rows int
	cols int
	path [][]int
)

func exist(board []string, word string) bool {
	rows = len(board)
	cols = len(board[0])
	path = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	dp := make([][]byte, rows)
	for i := range dp {
		dp[i] = make([]byte, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dp[i][j] = byte(board[i][j])
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dp[i][j] == word[0] && dfs(dp, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func dfs(dp [][]byte, word string, row int, col int, strIdx int) bool {
	if dp[row][col] != word[strIdx] {
		return false
	}
	if strIdx == len(word)-1 {
		return true
	}
	dp[row][col] = '0'
	for i := 0; i < 4; i++ {
		tempRow := row + path[i][0]
		tempCol := col + path[i][1]
		if !(tempRow < 0 || tempRow >= rows || tempCol < 0 || tempCol >= cols) {
			if dfs(dp, word, tempRow, tempCol, strIdx+1) {
				return true
			}
		}
	}
	dp[row][col] = word[strIdx]
	return false
}
