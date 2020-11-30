package main

import "fmt"

/**
*	结构只能比较是否相等、不能比较大小
*	相同类型的结构体才能进行比较、结构体是否相等不单与属性类型有关、还与属性的顺序有关。
*	可比较类型：int、float，bool、字符串、指针、数组等。
*	不可比较的类型：切片、map、函数等不可以比较。
 */
func main() {
	day5Test01()
}

func day5Test01() {
	sn1 := struct {
		name string
		age  int
	}{name: "w", age: 10}

	sn2 := struct {
		name string
		age  int
	}{name: "w", age: 10}
	fmt.Println(sn1 == sn2)

}

func day5Test02() {
	sm1 := struct {
		age int
		m   map[string]int
	}{age: 1, m: map[string]int{"a": 1}}

	sm2 := struct {
		age int
		m   map[string]int
	}{age: 1, m: map[string]int{"a": 1}}
	//fmt.Println(sm1==sm2)
	fmt.Println(sm1, sm2)
}
