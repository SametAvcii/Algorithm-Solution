package solution

import (
	"math"
)

func Solution(N int) int {
	a := []string{"1"}
	i := 0.0

	for int(math.Pow(2, i)) <= N {
		i++
	}
	i -= 1
	b := int(i)

	cur := int(math.Pow(2, i-1))
	max := N - int(math.Pow(2, i))

	for x := 0; x < b; x++ {
		if N-cur == 0 {
			for j := 0; j < b; j++ {
				a = append(a, "0")
			}
			break
		}
		if max-cur >= 0 {
			a = append(a, "1")
			max -= cur
		} else {
			a = append(a, "0")
		}
		i--
		cur = int(math.Pow(2, (i - 1)))
	}

	d := []int{}
	count := 0
	for _, i := range a {
		if i == "0" {
			count += 1
			d = append(d, count)
		} else {
			count = 0
		}
	}

	big := 0
	control := false
	for _, x := range d {
		if x > big {
			big = x
		} else {
			control = true
		}
	}
	if control == false {
		return 0
	}
	return big
}
