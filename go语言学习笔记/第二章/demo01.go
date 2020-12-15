package 第二章

import "fmt"

const (
	a = iota
	b
	c = 100
	d
	e = iota
	f
)

const (
	g         = iota
	h float64 = iota
	i
)

func Print() {
	fmt.Println(a, b, c, d, e, f)
	fmt.Println("----")
	fmt.Println(g, h, i)
}
