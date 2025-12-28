package main

func main() {

}

func calculateScore(instructions []string, values []int) int64 {
	res := 0
	i := 0
	visited := make([]bool, len(instructions))
	for i >= 0 && i < len(instructions) && !visited[i] {
		visited[i] = true
		if instructions[i] == "add" {
			res += values[i]
			i++
		} else {
			i = i + values[i]
		}
	}
	return int64(res)
}

var (
	path []int
	res  int
)

func maximumPossibleSize(nums []int) int {
	path = []int{}
	res = 0
	dfs(nums, 0)
	return res
}
func dfs(nums []int, idx int) {
	if idx == len(nums) {
		res = max(res, len(path))
	}
	if len(path) == 0 || nums[idx] >= path[len(path)-1] {
		path = append(path, nums[idx])
		dfs(nums, idx+1)
		path = path[:len(path)-1]
	}
	dfs(nums, idx+1)
}
