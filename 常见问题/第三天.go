package main

import "fmt"

func main() {
	fmt.Println(day3Test01())
	fmt.Println(day3Test02())
}

func day3Test01() []int {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	return s
}

func day3Test02() []int {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	return s
}
