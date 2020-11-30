package main

import "fmt"

func main() {
	day11Test01()
	day11Test02()

}
func day11Test01() {
	a := [1]int{1}
	b := []int{1}
	c := map[string]int{"a": 1}
	d := make(chan int, 1)
	fmt.Println(cap(a))
	fmt.Println(cap(b))
	//fmt.Println(cap(c)) cap函数不适用于map
	fmt.Println(c)
	fmt.Println(cap(d))

}
func day11Test02() {
	s := make(map[string]int)
	delete(s, "h")      //之不存在、没任何影响
	fmt.Println(s["h"]) //访问是可以的 0
}
