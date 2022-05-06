package solution

// you can also use imports, for example:

import "strings"
import "strconv"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(E string, L string) int {

	LastTime := strings.Split(L, ":")
	FirstTime := strings.Split(E, ":")
	IntLastTime := []int{}
	IntFirstTime := []int{}

	for _, i := range LastTime {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		IntLastTime = append(IntLastTime, j)
	}
	for _, i := range FirstTime {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		IntFirstTime = append(IntFirstTime, j)
	}

	TotalHours := IntLastTime[0] - IntFirstTime[0]
	TotalMinute := IntLastTime[1] - IntFirstTime[1]

	for TotalMinute/60 > 1 {
		TotalHours += 1
	}

	TotalCost := 0
	for i := 0; i < TotalHours+1; i++ {

		if i == 0 {
			TotalCost += 2
		} else if i == 1 {
			TotalCost += 3
		} else {
			TotalCost += 4
		}
	}
	if TotalMinute < 60 && TotalMinute > 0 {
		TotalCost += 4
	}
	return TotalCost

}
