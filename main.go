package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a) && i%2 == 0; i++ {
		fmt.Println(a[i])
	}
}
