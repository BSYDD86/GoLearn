package main

func main() {
	arr := []int{1, 2, 87, 87, 87, 2, 1}
	candy(arr)

}

func candy(ratings []int) int {
	need := make([]int, len(ratings))
	for i := range need {
		need[i] = 1
	}
	n := len(ratings)
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			need[i]++
		}
	}
	for j := n - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			need[j] = max(need[j], need[j+1]+1)
		}
	}
	res := 0
	for _, cnt := range need {
		res += cnt
	}
	return res
}
