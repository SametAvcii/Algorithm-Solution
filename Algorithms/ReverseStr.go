package main

import (
	"fmt"
	"strings"
)

func ReverseStr(n string) string {
	x := strings.Split(n, "")
	l := len(x) - 1
	i := 0
	for l > i {
		x[i], x[l] = x[l], x[i]
		i++
		l--
	}
	a := ""
	for _, j := range x {
		a += j

	}
	return a

}

func main() {
	n := "abcdefg"
	fmt.Println(ReverseStr(n))
}
