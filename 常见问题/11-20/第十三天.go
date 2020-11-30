package main

import "fmt"

func day13Test01(i int) {
	fmt.Println(i)
}
func main() {
	i := 5
	defer day13Test01(i)
	defer func() {
		day13Test01(i)
	}()
	i = i + 10
}
