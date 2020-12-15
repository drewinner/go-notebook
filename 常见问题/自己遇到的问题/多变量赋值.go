package main

import "fmt"

//go build -o test 多变量赋值.go
//go tool objdump "main\.main" test
func main() {
	var a float64 = 20

	switch a {
	case 20:
		fmt.Println("yes")
	default:
		fmt.Println("default")
	}
}
