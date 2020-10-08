package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	strArr := []string{}
	for i := 0; i< math.MaxUint8; i++ {
		if s, ok := NextString(i); ok {
			strArr = append(strArr, s)
		}
	}
	fmt.Println(strings.Join(strArr, ""))
	 
}
