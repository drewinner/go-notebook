package main

import "fmt"

func main() {

}

func day4Test01() {
	list := new([]int)
	//list = append(*list, 1)
	//list = append(list, 1)
	fmt.Println(list)
}

func day4Test02() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//s1 = append(s1, s2) //错误、第二个参数是可变参数
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

var (
	//size := 1024 //短命名
	size     = 1024 //短命名
	max_size = size * 2
)

func day4Test03() {
	//size int  := 1024 //短命名、不能有数据类型
	//var a int = 15 //这个可以
	fmt.Println(size, max_size)
}
