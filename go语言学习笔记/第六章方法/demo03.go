package 第六章方法

import "fmt"

type X struct {
}

func (x *X) Test() {
	fmt.Println("hi !", x)
}
