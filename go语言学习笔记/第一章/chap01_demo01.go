package 第一章

import "fmt"

type user struct {
	Name string
	Age  int
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v\n", u)
}

type manager struct {
	user
	title string
}
