package main

import "fmt"

type structInterface interface {
	StructDemo01()
	StructDemo02()
}

type StructS struct {
	structInterface
}

func GetStructS(sI structInterface) structInterface {
	return &StructS{structInterface: sI}
}

func (s StructS) StructDemo01() {
	fmt.Println("exec demo01 ...")
	s.structInterface.StructDemo01()
}

//实现
type S struct {
	StructS
}

func (s *S) StructDemo01() {
	fmt.Println("StructDemo01..")
}

func main() {
	var s S
	s1 := GetStructS(&s)
	s1.StructDemo01()
}
