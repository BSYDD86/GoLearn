package main

func main() {
	const mx = 101
}

func dfs(nums []int, idx int, path int) int {
	if path < 0 {
		return 0
	} else if path == 0 {
		return 1
	}
	if idx == len(nums) {
		return 0
	}
	res := 0
	for i := idx; i < len(nums); i++ {
		res += dfs(nums, i, path-nums[i])
	}
	return res
}
