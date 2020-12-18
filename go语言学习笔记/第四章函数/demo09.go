package 第四章函数

import "fmt"

func Test09() []func() {
	var s []func()
	for i := 0; i < 2; i++ {
		s = append(s, func() {
			fmt.Println(&i, i)
		})
	}
	return s
}

//参数和返回值=函数类型
func Test0901(x int) (func(), func()) {
	return func() {
			fmt.Println(x)
			x += 10
		}, func() {
			fmt.Println(x)
		}
}
