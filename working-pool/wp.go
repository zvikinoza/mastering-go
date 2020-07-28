package main

import (
	"fmt"
)

func work(job int) int {
	return job * 1000
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker ", id, " started job", job)
		results <- work(job)
		fmt.Println("worker ", id, " ended job", job)
	}
}

func main() {
	const numJobs, numWorkers = 500000, 300000
	jobs, results := make(chan int, numJobs), make(chan int, numJobs)

	for wid := 0; wid < numWorkers; wid++ {
		go worker(wid, jobs, results)
	}

	for job := 0; job < numJobs; job++ {
		jobs <- job
	}
	//close(jobs)

	for res := 0; res < numJobs; res++ {
		fmt.Println("result of", res, "=", <-results)
	}
}
