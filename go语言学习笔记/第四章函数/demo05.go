package 第四章函数

import "fmt"

func Test05(s []string) {
	fmt.Printf("pointer:%p target:%v\n", &s, s)
	s[1] = "bbb"
	s[2] = "ccc"
	s = append(s, "cc")
}
