package main

import (
	"fmt"
	"reflect"
)

var (
	res []string
)

func main() {
	fmt.Println(reflect.TypeOf('0'))
	fmt.Println(reflect.TypeOf(0))
	res = make([]string, 0)
	path := ""
	dfs(0, 2, path)
	for i, num := range res {
		pre := 0
		for pre < len(num)-1 {
			if num[pre] == '0' {

				pre++
			} else {
				break
			}
		}
		res[i] = num[pre:len(num)]
	}
	fmt.Println(res)
}

func dfs(curIdx, n int, path string) {
	if curIdx == n {
		temp := path
		res = append(res, temp)
		return
	}
	for i := 0; i < 10; i++ {
		path = path + string(i+'0')
		dfs(curIdx+1, n, path)
		path = path[:len(path)-1]
	}
}
