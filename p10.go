package main

func main() {
	s := []int{1,2,3,4,5}

	s = insert(s, []int{-3, -2}, 0)
}

func insert(s, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}

func insert2(s, new []int, index int) []int {
	result := make([]int, len(s) + len(new))
	position := copy(result, s[:index])
	position += copy(result[position:], new)
	copy(result[position:], s[index:])
	return result
}

