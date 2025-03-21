package main

func main() {
	//str := "3[a2[c]]"
	//fmt.Println(decodeString(str))
}

//	func canFinish(numCourses int, prerequisites [][]int) bool {
//		g := make([]int, numCourses)
//		mapPreReques := make([][]int, numCourses)
//		for _, pre := range prerequisites {
//			mapPreReques[pre[1]] = append(mapPreReques[pre[1]], pre[0])
//		}
//		//have chain
//		var dfs func(int) bool
//		dfs = func(x int) bool {
//			g[x] = 1
//			for _, y := range mapPreReques[x] {
//				if g[y] == 1 || (g[y] == 0 && dfs(y)) {
//					return true
//				}
//			}
//			g[x] = 2
//			return false
//		}
//	}
func countPrimes(n int) int {
	res := 0
	for i := 2; i < n; i++ {
		if isPrimes(res) {
			res++
		}
	}
	return res
}

func isPrimes(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isValid(s string) bool {
	st := []byte{}
	for _, c := range s {
		tempC := byte(c)
		if tempC == '(' || tempC == '[' || tempC == '{' {
			st = append(st, tempC)
		} else {
			if len(st) == 0 {
				return false
			}
			tempC = reverse(tempC)
			if st[len(st)-1] != tempC {
				return false
			} else {
				st = st[:len(st)-1]
			}
		}
	}
	return len(st) == 0
}

func reverse(c byte) byte {
	if c == '(' {
		return ')'
	} else if c == '[' {
		return ']'
	} else {
		return '}'
	}
}
