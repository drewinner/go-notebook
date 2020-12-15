package 第二章

import "fmt"

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func Print04() {
	fmt.Println(read, write, exec)
}
