package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var opsAtomic atomic.Uint64
	var ops int

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				opsAtomic.Add(1)
				ops++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("OpsAtomic", opsAtomic.Load())
	fmt.Println("Ops", ops)
}
