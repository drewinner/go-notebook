package main

import "fmt"

//https://www.cnblogs.com/zsy/p/5370052.html
const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

//遇到新的const从头开始计数、也可以解释称iota是在const块内部索引
const (
	a = iota
	b
	_
	d
)

func main() {
	fmt.Println(x, y, z, k, p) //0,2,zz,zz,5
	fmt.Println(a, b, d)
}

func day7Test01() {
	//fmt.Println(iota) //iota只能在常量表达是中使用
}

//nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量
func day7Test02() {
	//var x = nil
	var y interface{} = nil
	//var z string = nil
	var e error = nil
	fmt.Println(y, e)
}
