package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]

	fmt.Println(len(a), cap(a)) //0,3
	fmt.Println(len(b), cap(b)) //2,3
	fmt.Println(len(c), cap(c)) //1,3 这里我错了。。
}
