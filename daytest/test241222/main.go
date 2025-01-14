package main

import "fmt"

func main() {
	s := "()[]{}"
	isValid(s)
}

func isValid(s string) bool {
	//nil to solve
	if len(s)%2 == 1 {
		return false
	}
	st := []rune{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			t := convert(c)
			st = append(st, t)
		} else {
			if len(st) == 0 || st[len(st)-1] != c {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	fmt.Println(1 + 1)
	return len(st) == 0
}

func convert(c rune) rune {
	if c == '(' {
		return ')'
	}
	if c == '[' {
		return '}'
	}
	if c == '{' {
		return '}'
	}
	return '1'
}
