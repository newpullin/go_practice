package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	done := process(quit)
	timeout := time.After(10 * time.Millisecond)

	select {
	case d := <-done:
		fmt.Println(d)
	case <-timeout:
		fmt.Println("Time out!")
	}
}

func process(quit <-chan struct{}) chan string {
	done := make(chan string)
	go func() {
		go func() {
			time.Sleep(10 * time.Second)

			done <- "Complete!"
		}()

		<-quit
		return
	}()

	return done
}
