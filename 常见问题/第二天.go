package main

import "fmt"

// 这是新手常会犯的错误写法，for range 循环的时候会创建每个元素的副本，
// 而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
// 所以最后 map 中的所有元素的值都是变量 val 的地址，
// 因为最后 val 被赋值为3，所有输出都是3.
func main() {

	slice := []string{"a", "b", "c", "d"}
	m := make(map[int]*string)
	for key, val := range slice {
		v := val
		m[key] = &v
		//m[key] = &val //错误的下发
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
	fmt.Println("----")
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	for _, v := range m1 {
		fmt.Println(&v)
	}
}
