package main

import "fmt"

//所谓引用类型特指：slice、map、channel三种预定义类型

func main() {
	m := mkmap()
	fmt.Println(m["a"])
	s := mkslice()
	fmt.Println(s[0])
}
func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}
