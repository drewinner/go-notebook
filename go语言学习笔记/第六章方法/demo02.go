package 第六章方法

import "fmt"

func (n N) value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("v:%p,%v\n", n, *n)
}
