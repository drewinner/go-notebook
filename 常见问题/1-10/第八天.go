package main

import "fmt"

func main() {
	h := day8Test01
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	//i := day8Test02()
	//编译失败。考点：类型选择，类型选择的语法形如：i.(type)，
	// 其中 i 是接口，type 是固定关键字，需要注意的是，
	// 只有接口类型才可以使用类型选择。看下关于接口的文章。
	//switch i.(type) {
	//case int:
	//	println("int")
	//case string:
	//	println("string")
	//case interface{}:
	//	println("interface")
	//default:
	//	println("unknown")
	//}
}

func day8Test01() []string {
	return nil
}

func day8Test02() int {
	return 1
}
