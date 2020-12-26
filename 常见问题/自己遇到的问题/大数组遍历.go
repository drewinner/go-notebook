package main

import (
	"fmt"
	"time"
)

func main() {
	s := time.Now()
	var arr = [10002400]int{1, 1, 1}
	fmt.Println(time.Now().Sub(s))
	s1 := time.Now()
	for i, n := range arr {
		//just ignore i and n for simplify the example
		_ = i
		_ = n
	}
	fmt.Println(time.Now().Sub(s1))
}
