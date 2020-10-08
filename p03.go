package main

import "fmt"

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}


func main() {
	f()
	fplus := func(x,y int) int {
		return x+y
	}
	fmt.Println(fplus(3,5))
}
