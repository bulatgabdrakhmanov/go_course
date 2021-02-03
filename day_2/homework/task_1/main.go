package main

import "fmt"

func createArray(N int) []int {
	arr := make([]int, 0, N)

	for i := 1; i <= N; i++ {
		arr = append(arr, i)
	}

	return arr
}

func shift(arr []int, steps int) []int {
	arrLen := len(arr)
	steps %= arrLen

	tmp := make([]int, arrLen)
	copy(tmp, arr)
	for i := 0; i < arrLen; i++ {
		arr[(i+steps)%arrLen] = tmp[i]
	}

	return arr
}

func main() {
	arr := createArray(6)
	fmt.Println(shift(arr, 2))
}
