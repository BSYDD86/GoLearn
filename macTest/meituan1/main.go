package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		cnt := 0
		res := make([]byte, 0)
		for i := range str {
			c := str[i]
			if c >= '0' && c <= '9' {
				atoi, _ := strconv.Atoi(str[i : i+1])
				if cnt == 0 {
					cnt = atoi
				} else {
					cnt = cnt*10 + atoi
				}
			} else {
				if len(res) != 0 {
					cnt %= len(res)
					cnt = min(cnt, len(res))
					if cnt != 0 {
						temp := res[:cnt]
						res = res[cnt:]
						res = append(res, temp...)
						cnt = 0
					}
				}

				if c == 'R' {
					l := 0
					r := len(res) - 1
					for l < r {
						res[l], res[r] = res[r], res[l]
						l++
						r--
					}
				} else {
					res = append(res, c)
				}
			}

		}
		fmt.Println(string(res))
	}
}
