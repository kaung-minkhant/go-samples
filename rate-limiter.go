package main

import (
	"fmt"
	"time"
)

func main() {
	// requests := make(chan int, 5)
	// for j := 0; j < 5; j++ {
	// 	requests <- j
	// }
	// close(requests)
	//
	// limiter := time.NewTicker(1 * time.Second)
	// for req := range requests {
	// 	<-limiter.C
	// 	fmt.Println("Request:", req, "Time:", time.Now())
	// }

	burstRequest := make(chan int, 5)
	for j := 0; j < 5; j++ {
		burstRequest <- j
	}
	close(burstRequest)

	const burstLimit = 5
	burstLimiter := make(chan time.Time, burstLimit)
	for j := 0; j < burstLimit; j++ {
		burstLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(1 * time.Second) {
			burstLimiter <- t
		}
	}()

	for req := range burstRequest {
		<-burstLimiter
		fmt.Println("Request:", req, "Time:", time.Now())
	}

}
