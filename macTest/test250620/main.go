package main

import "fmt"

func main() {
	str := "NWSE"
	k := 1
	fmt.Println(maxDistance(str, k))
}

func maxDistance(s string, k int) int {
	arr := make([]int, 4)
	for _, c := range s {
		if c == 'N' {
			arr[0]++
		} else if c == 'S' {
			arr[1]++
		} else if c == 'E' {
			arr[2]++
		} else if c == 'W' {
			arr[3]++
		}
	}
	cnt := max(arr[0], arr[1]) + max(arr[2], arr[3])
	temp1 := arr[0] - arr[1]
	temp2 := arr[2] - arr[3]
	if temp1 < 0 {
		temp1 = -temp1
	}
	if temp2 < 0 {
		temp2 = -temp2
	}
	res := temp1 + temp2
	cnt = min(cnt, k)
	res += cnt * 2
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumLength(s string) int {
	i := 0
	j := len(s) - 1
	for i < j && s[i] == s[j] {
		temp := s[i]
		for i <= j && s[i] == temp {
			i++
		}
		for i <= j && s[j] == temp {
			j--
		}
	}
	return -1
}
