package 第四章函数

import "fmt"

func Test04(p *int) {
	go func() {
		fmt.Println(p)
	}()
}
