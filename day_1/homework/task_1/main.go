package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNaturalSquareNumbers(numsCount int) []int {
	var result []int

	for i := 0; i < numsCount; i++ {
		randNum := rand.Intn(10000) + 1
		result = append(result, randNum*randNum)
	}

	return result
}

func main() {
	fmt.Println(generateRandomNaturalSquareNumbers(10000))
}
