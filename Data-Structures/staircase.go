package main

import "fmt"

func staircase(n int32) {
	for i := 1; i < int(n+1); i++ {
		for j := 0; j < int(n)-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}

		fmt.Println()
	}

}
func main() {
	staircase(5)
}
