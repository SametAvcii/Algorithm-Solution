package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int) string {
	Letter := ""
	Bigger := 0
	Smaller := 0
	if A < B {
		Bigger = B
		Smaller = A
	} else {
		Bigger = A
		Smaller = B
	}
	if Bigger/Smaller > 2 {

		if Bigger == A {
			for Bigger != 0 {
				Letter += "aa"
				Bigger -= 2

				if Smaller != 0 {
					Letter += "b"
					Smaller -= 1
				}
			}
		} else {
			for Bigger != 0 {
				Letter += "bb"
				Bigger -= 2

				if Smaller != 0 {
					Letter += "a"
					Smaller -= 1
				}
			}
		}
	} else if Bigger/Smaller < 2 {
		if Bigger == A {
			for Bigger != 0 {
				Letter += "a"
				Bigger -= 1
				if Smaller != 0 {
					Letter += "b"
					Smaller -= 1
				}
			}
		} else {
			for Bigger != 0 {
				Letter += "b"
				Bigger -= 1
				if Smaller != 0 {
					Letter += "a"
					Smaller -= 1
				}
			}
		}
	}
	return Letter
}


