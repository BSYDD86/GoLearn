package main

import "fmt"

func main() {
	k := int64(10)
	arr := []int{0, 1, 0, 1}
	fmt.Println(byte(kthCharacter(k, arr)))
}
func kthCharacter(k int64, operations []int) byte {
	cnt := 0
	for 1<<cnt-1 < k {
		cnt++
	}
	temp := 0
	for i := 0; i < cnt-1; i++ {
		temp += operations[i]
	}
	res := 0
	for k > 0 {
		if k&1 == 0 {
			res++
		}
		k >>= 1
	}

	res = min(res, temp)
	res %= 26
	return byte('a' + res)
}
