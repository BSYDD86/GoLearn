package temp

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	iMax := 0
	iMin := 0
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		if temp < 0 {
			iMin += temp
		} else {
			iMax += temp
		}
	}
	fmt.Println(iMax - iMin)
}
