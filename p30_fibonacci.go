package main

import (
	"fmt"
	"time"
)

const LIMIT = 1001
const PROCESS_NUM = 16

var memo [LIMIT]uint64

func main() {
	jobs := make(chan int, LIMIT)
	results := make(chan uint64, LIMIT)
	startTime := time.Now()

	for i := 0; i < PROCESS_NUM; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < LIMIT; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < LIMIT; j++ {
		fmt.Printf("%d = %d\n", j, <-results)
	}
	elapsedTime := time.Since(startTime)
	for j := 0; j < LIMIT; j++ {
		fmt.Printf("%d = %d\n", j, memo[j])
	}

	fmt.Printf("실행시간 : %s\n", elapsedTime)

}

func worker(jobs <-chan int, results chan<- uint64) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) uint64 {
	if memo[n] != 0 {
		return memo[n]
	}
	if n <= 1 {
		return uint64(n)
	}
	result := fib(n-1) + fib(n-2)
	if memo[n] == 0 {
		memo[n] = result
	}
	return result
}
