package main

import (
	"fmt"
	"strings"
)

func FindIntersection(strArr []string) string {

	strArr1 := strings.Split(strArr[0], ", ")
	strArr2 := strings.Split(strArr[1], ", ")
	var strArr3 []string
	for _, i := range strArr1 {
		for _, x := range strArr2 {

			if i == x {
				strArr3 = append(strArr3, x)
			}
		}
	}
	val := strings.Join(strArr3, ",")
	return val

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FindIntersection(readline()))

}
