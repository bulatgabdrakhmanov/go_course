package main

import (
	"fmt"
)

func printChan(channels ...<-chan int) {
	for _, ch := range channels {
		for item := range ch {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}
}

// Заполняет канал натуральными числами от 1 до n и возвращает этот канал
func generateChanWithNaturalNumbers(n int) chan int {
	ch := make(chan int, n)

	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

// Сливает один или большее количество каналов
func merge(channels ...<-chan int) chan int {
	bufferSize := 0

	for _, ch := range channels {
		bufferSize += cap(ch)
	}

	result := make(chan int, bufferSize)
	for _, ch := range channels {
		for item := range ch {
			result <- item
		}
	}

	close(result)
	return result
}

func main() {
	ch1 := generateChanWithNaturalNumbers(50) // chan(1, 2, ..., 50)
	ch2 := generateChanWithNaturalNumbers(30) // chan(1, 2, ..., 30)

	printChan(merge(ch2, ch1)) // chan(1, 2, ..., 30, 1, 2, ..., 50)
}
