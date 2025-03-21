package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "2"
	num2 := "3"
	fmt.Println(multiply(num1, num2))

}

func multiply(num1 string, num2 string) string {
	if num1 == "" || num1 == "0" {
		return "0"
	}
	if num2 == "" || num2 == "0" {
		return "0"
	}
	dp := make([]int, len(num1)+len(num2))
	m := len(num1)
	n := len(num2)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			v1, _ := strconv.Atoi(num1[i : i+1])
			v2, _ := strconv.Atoi(num2[j : j+1])
			idx := i + j + 1
			temp := v1*v2 + dp[idx]
			for temp > 0 {
				dp[idx] = temp % 10
				temp /= 10
				if temp > 0 {
					idx--
					temp += dp[idx]
				}
			}
		}
	}
	res := ""
	idx := 0
	for len(res) == 0 && dp[idx] == 0 {
		idx++
	}
	for idx < m+n {
		v := strconv.Itoa(dp[idx])
		res = res + v
		idx++
	}
	return res
}
