package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	binaryN := strconv.FormatInt(int64(N), 2)

	consecutiveZerosCounter, consecutiveZerosMax := 0, 0
	for _, e := range binaryN {
		if e-'0' == 0 {
			consecutiveZerosCounter += 1
			if consecutiveZerosCounter > consecutiveZerosMax {
				consecutiveZerosMax = consecutiveZerosCounter
			}
		} else {
			consecutiveZerosCounter = 0
		}
	}

	return consecutiveZerosMax
}

func main() {
	fmt.Println(Solution(48))
	fmt.Println(Solution(23))
	fmt.Println(Solution(1024))
	fmt.Println(Solution(0))
}
