package main

import "fmt"

var aa bool = true

//defer 关键字后面的函数或者方法想要执行必须先注册
func main() {
	defer func() {
		fmt.Println("1")
	}()
	if aa == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
