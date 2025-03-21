package main

import "fmt"

func main() {
	arr := [][]int{{1}, {2}, {3}, {}}
	fmt.Println(canVisitAllRooms(arr))
}

func backspaceCompare(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sIdx := len(s) - 1
	tIdx := len(t) - 1
	for sIdx >= 0 && tIdx >= 0 {
		for sIdx >= 0 && s[sIdx] == '#' {
			sIdx -= 2
		}
		for tIdx >= 0 && t[tIdx] == '#' {
			tIdx -= 2
		}
		if tIdx < 0 || sIdx < 0 {
			break
		}
		if s[sIdx] == t[tIdx] {
			sIdx--
			tIdx--
		} else {
			return false
		}
	}
	return tIdx < 0 && sIdx < 0
}

func canVisitAllRooms(rooms [][]int) bool {
	res := 1
	m := map[int]bool{}
	m[0] = true
	st := []int{0}
	for len(st) > 0 {
		idx := st[len(st)-1]
		st = st[:len(st)-1]

		for _, n := range rooms[idx] {
			if !m[n] {
				m[n] = true
				res++
				st = append(st, n)
			}
		}

	}
	return res == len(rooms)
}
