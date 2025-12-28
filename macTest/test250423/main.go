package main

import "fmt"

func main() {
	group := countLargestGroup(46)
	fmt.Println(group)
}

func countLargestGroup(n int) int {
	m := map[int]int{}
	res := 0
	cnt := 0
	for i := 1; i <= n; i++ {
		sum := sumEveryDigit(i, 10)
		fmt.Println(i, sum)
		m[sum]++
		if m[sum] == res {
			cnt++
		} else {
			res = m[sum]
			cnt = 1
		}
	}
	fmt.Println(m)
	temp := 0
	for _, v := range m {
		temp += v
	}
	fmt.Println(temp)
	return cnt
}
func sumEveryDigit(num int, radix int) int {
	res := 0
	for num > 0 {
		res += num % radix
		num /= radix
	}
	return res
}
