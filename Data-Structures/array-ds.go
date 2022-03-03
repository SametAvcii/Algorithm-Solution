package main

import "fmt"

func reverse(a []int) []int {

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {

		a[i], a[j] = a[j], a[i]

	}
	return a
}
func main() {
	a := []int{5, 7, 8, 9, 0, 6, 432, 4, 5}
	fmt.Println(reverse(a))

}
