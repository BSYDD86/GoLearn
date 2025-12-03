package test250929

func main() {

}

func largestTriangleArea(points [][]int) float64 {
	res := float64(0)
	for i := 0; i+2 < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				x1 := points[i][0]
				x2 := points[j][0]
				x3 := points[k][0]
				y1 := points[i][0]
				y2 := points[j][1]
				y3 := points[k][2]
				res = max(res, getArea(x1, y1, x2, y2, x3, y3))
			}
		}
	}
	return res
}

func max(res float64, area float64) float64 {
	return max(res, area)
}

func getArea(x1, y1, x2, y2, x3, y3 int) float64 {
	temp := float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)) * 0.5
	if temp < 0 {
		temp = -temp
	}
	return temp
}
