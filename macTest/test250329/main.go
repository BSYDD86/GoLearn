package main

import "fmt"

func main() {
	arr := []int{123, 456, 789, 13, 25}
	L, R := 0, len(arr)-1
	// 假设最大数是三位数，这里digit设为3，实际应用中应准确获取最大数的位数
	digit := 3
	radixSort(arr, L, R, digit)
	fmt.Println(arr)
}

func getDigit(num, d int) int {
	divisor := 1
	for i := 1; i < d; i++ {
		divisor *= 10
	}
	return (num / divisor) % 10
}

func radixSort(arr []int, L, R, digit int) {
	radix := 10
	bucket := make([]int, R-L+1)
	for d := 1; d <= digit; d++ {
		count := make([]int, radix)
		// 统计当前位上每个数字出现的次数
		for i := L; i <= R; i++ {
			j := getDigit(arr[i], d)
			count[j]++
		}
		// 对count数组进行累加
		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}
		// 将元素按当前位的数值分配到桶中
		for i := R; i >= L; i-- {
			j := getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}
		// 将桶中的元素复制回原数组
		for i, j := L, 0; i <= R; i, j = i+1, j+1 {
			arr[i] = bucket[j]
		}
	}
}

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := range dp {
		dp[i] = make([]int, numRows)
	}
	for i := 0; i < numRows; i++ {
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp
}
