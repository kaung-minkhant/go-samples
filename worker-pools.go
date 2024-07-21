package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(1 * time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numOfJobs = 5
	const numOfWorkers = 3
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for j := range numOfWorkers {
		go worker(j, jobs, results)
	}

	for j := range numOfJobs {
		jobs <- j
	}
	close(jobs)

	for j := 0; j < numOfJobs; j++ {
		<-results
	}
}
