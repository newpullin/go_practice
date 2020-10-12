package main

import "fmt"

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r *rect) resize(w, h float64) {
	r.width += w
	r.height += h
}

func main() {
	r := rect{3, 4}
	fmt.Println("area : ", r.area())
	r.resize(10, 10)
	fmt.Println("area :", r.area())

	/*
		area() 메서드의 함수 표현식
		서명: func(rect) float64
	*/
	areaFn := rect.area
	resizeFn := (*rect).resize

	fmt.Println("Area :", areaFn(r))
	resizeFn(&r, -10, -10)
	fmt.Println("area :", areaFn(r))
}
