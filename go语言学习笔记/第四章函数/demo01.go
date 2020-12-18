package 第四章函数

import "fmt"

/**
*	函数属于第一类对象、具备相同签名（参数和返回值）视作同一类型
 */
type f func(s string) (str string)

func test01() {
	fmt.Println("hello world")
}
func Exec01(f func()) {
	f()
}
func test02(str string) string {
	return str
}

func Exec02(instr string, fn f) string {
	return fn(instr)
}
