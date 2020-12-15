package 第一章

import "fmt"

type user02 struct {
	name string
	age  int
}

func (u user02) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {
	Print()
}
