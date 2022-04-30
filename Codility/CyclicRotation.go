package Codility

func Solution(A []int, K int) []int {

	len := len(A)

	for j := 0; j < K; j++ {
		var B []int
		B = append(B, A...)

		last := B[len-1]
		for x, i := range B {

			if x == len-1 {
				A[0] = last
			} else {
				A[x+1] = i
			}
		}
	}
	return A
}
