package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createArray(N int) []int {
	arr := make([]int, 0, N)

	rand.Seed(int64(time.Now().Unix()))

	for i := 1; i <= N; i++ {
		arr = append(arr, rand.Intn(N/2))
	}

	return arr
}

func notUniqueItemsCount(arr []int) int {
	existedItemCount := make(map[int]int)
	uniqueItemsCount := len(arr)

	for _, e := range arr {
		existedItemCount[e]++

		if existedItemCount[e] >= 2 {
			if existedItemCount[e] == 2 {
				uniqueItemsCount--
			}
			uniqueItemsCount--
		}
	}

	return uniqueItemsCount
}

func main() {
	randArr := createArray(80)
	fmt.Println(randArr)
	fmt.Println(notUniqueItemsCount(randArr))
}
