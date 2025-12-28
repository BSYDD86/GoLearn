package main

func main() {
	status := []int{1, 0, 1, 0}
	candies := []int{7, 5, 4, 100}
	keys := [][]int{{}, {}, {1}, {}}
	containedBoxes := [][]int{{1, 2}, {3}, {}, {}}
	initialBoxes := []int{0}
	// 这里假设已经有正确实现的maxCandies函数，可将上述数据传入进行计算
	result := maxCandies(status, candies, keys, containedBoxes, initialBoxes)
	println(result)
}

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	box := []int{}
	for _, num := range initialBoxes {
		box = append(box, num)
	}
	key := map[int]bool{}
	change := true
	res := 0
	for len(box) > 0 && change {
		deleted := []int{}
		change = false
		tempSize := len(box)
		for i := 0; i < tempSize; i++ {
			boxIdx := box[i]
			if status[boxIdx] == 1 {
				deleted = append(deleted, i)
				res += candies[boxIdx]
				change = true
				for _, v := range keys[boxIdx] {
					key[v] = true
				}
				for _, v := range containedBoxes[boxIdx] {
					box = append(box, v)
				}
			} else if status[boxIdx] == 0 && key[boxIdx] {
				deleted = append(deleted, i)
				res += candies[boxIdx]
				change = true
				for _, v := range keys[boxIdx] {
					key[v] = true
				}
				for _, v := range containedBoxes[boxIdx] {
					box = append(box, v)
				}
			}
		}
		for i := len(deleted) - 1; i >= 0; i-- {
			deletedIdx := deleted[i]
			box = append(box[0:deletedIdx], box[deletedIdx+1:]...)
		}
	}
	return res
}
