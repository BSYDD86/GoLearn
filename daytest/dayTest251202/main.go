package main

func main() {
	point := [][]int{
		{1, 0},
		{2, 0},
		{3, 0},
		{2, 2},
		{3, 2},
	}
	countTrapezoids(point)

}

func countTrapezoids(points [][]int) int {
	const mod = 1e9 + 7
	m := make(map[int]map[int]int)
	for i, p := range points {
		x := p[0]
		y := p[1]
		for j := 0; j < i; j++ {
			p2 := points[j]
			x2 := p2[0]
			y2 := p2[1]
			if y == y2 {
				delta := x2 - x
				if delta < 0 {
					delta = -delta
				}
				if m[delta] == nil {
					m[delta] = make(map[int]int)
				}
				m[delta][y2]++
			}
		}
	}
	res := 0
	for k1 := 0; k1 < len(m); k1++ {
		v1 := m[k1]
		if len(v1) == 1 {
			continue
		}
		pre := 0
		for _, v := range v1 {
			res += pre * v
			pre += v
		}
	}
	return (res / 2) % mod
}
