package main
import (
	"fmt"
	"sort"
)

func Reptetive(a []int) []int {
	sort.Ints(a)
	b := len(a) / 2
	rep := []int{}

	for x, i := range a {
		if x == a[len(a)-1] {
			break
		}
		var c []int
		if i < b {
			c = append(a[:x])
		} else {
			c = append(a[x+1:])
		}

		for _, z := range c {
			if i == z {
				rep = append(rep, i)
			}
			rep = append(rep)
		}

	}
	if a[len(a)-1] == a[len(a)-2] {
		rep = append(rep, a[len(a)-1])
	}

	return rep

}
//This solution has O(logn) time Complexity

func main(){
	num:={2,3,4,5,5,4,7,8,9,9}
	fmt.Println(Reptetive(num))
}