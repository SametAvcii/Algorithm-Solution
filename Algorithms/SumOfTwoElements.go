package main

import "fmt"

func Sum(x []int) int {
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}
func Piramit(a, b, c, d int) int {
	x := []int{a, b, c, d}
	z := Sum(x) / 2
	for _, i := range x {
		for _, x := range x {
			if i+x == z {
				return z
			}
		}
	}
	return -1

}
func main() {

	fmt.Println(Piramit(1, 2, 10, 20))

}
