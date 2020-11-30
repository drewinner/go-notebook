package main

import "fmt"

//考点：可变函数
func main() {
	i := []int{5, 6, 7}
	day9Test01(i...)
	fmt.Println(i[0])
}
func day9Test01(num ...int) {
	num[0] = 18
}
