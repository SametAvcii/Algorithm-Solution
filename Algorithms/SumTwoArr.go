package main

import "fmt"

func Sum(l1, l2 []int) []int {
	l3 := []int{}
	for i, _ := range l1 {
		sum := l1[i] + l2[i]
		l3 = append(l3, sum)
	}
	return l3
}

func SumTwoList(l1, l2 []int) []int {
	l3 := []int{}
	if len(l1) == len(l2) {
		l3 = Sum(l1, l2)
	} else if len(l1) > len(l2) {
		i := 0
		for len(l1) != len(l2) {
			l2 = append(l2, l2[i])
			i++
		}
		l3 = Sum(l1, l2)
	} else {
		l3 = Sum(l1, l2)
	}
	return l3
}
func main() {

	l1 := []int{1, 2, 3, 1, 2}
	l2 := []int{8, 7, 6, 5}
	fmt.Println(SumTwoList(l1, l2))

}
