package 第二章

import "fmt"

type color byte //自定义类型

const (
	black color = iota
	red
	blue
)

func Test(c color) {
	fmt.Println(c)
}
