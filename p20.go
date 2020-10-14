package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main함수 시작", time.Now())

	done := make(chan bool)
	go long(done)
	go short(done)

	<-done
	<-done

	fmt.Println("main함수 종료", time.Now())
}

func long(done chan bool) {
	fmt.Println("long함수 시작 ", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("long함수 종료", time.Now())
	done <- true
}

func short(done chan bool) {
	fmt.Println("short함수 시작 ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short함수 종료", time.Now())
	done <- true
}
