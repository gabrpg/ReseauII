package arraygenerate

import "math/rand"

func ArrayGenerate(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(10000) + 1
	}
	return arr
}
