package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	d := NewDisJointSet(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			strA := strs[i]
			strB := strs[j]
			cnt := 0
			for i := 0; i < n; i++ {
				if strA != strB {
					cnt++
				}
				if cnt != 0 && cnt != 2 {
					continue
				}
				//fa := d.find(a)
				//fb := d.find(b)
				//if fa == fb {
				//	continue
				//}
				//d.union(a, b)
			}
		}
	}
	return d.sets
}

type DisJointSet struct {
	sets   int
	father []int
	size   []int
}

func NewDisJointSet(n int) *DisJointSet {
	d := &DisJointSet{
		sets:   n,
		father: make([]int, n),
		size:   make([]int, n),
	}
	for i := range d.father {
		d.father[i] = i
		d.size[i] = 1
	}
	return d
}

//
//func (d *DisJointSet) find(a) int {
//	if d.father[a] != a {
//		d.father[a] = find(d.father[a])
//	}
//	return d.father[a]
//}
//
//func (d *DisJointSet) isSameSet(a, b int) bool {
//	return d.father[a] == d.father[b]
//}
//
//func (d *DisJointSet) union(a, b int) {
//	fa := d.find(a)
//	fb := d.find(b)
//	if fa == fb {
//		return
//	}
//	if d.size(fa) > d.size(fb) {
//		fa, fb = fb, fa
//	}
//	d.father(fa) = fb
//	d.size(fb) += d.size(fa)
//	d.sets--
//}
