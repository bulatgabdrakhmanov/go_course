package main

import (
	"fmt"
	"time"
)

// Sleep реализованный на каналах
func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	start := time.Now().Unix()
	sleep(10000 * time.Millisecond)
	fmt.Println(time.Now().Unix() - start) // 10
}
