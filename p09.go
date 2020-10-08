package main

import "fmt"

func main() {
	numbers := []int{3, 4, 5, 6, 7, 8, 6, 8 ,32, 4}

	for _, value := range numbers {
		fmt.Print(value)
		fmt.Print(" ")
	}

	for i:= 0; i < len(numbers)-1; i++ {
		for k := 0; k < len(numbers)-i-1; k++ {
			if numbers[k] > numbers[k+1] {
				temp := numbers[k]
				numbers[k] = numbers[k+1]
				numbers[k+1] = temp
			}
		}
	}
	fmt.Println("")
	for _, value := range numbers {
		fmt.Print(value)
		fmt.Print(" ")
	}
}