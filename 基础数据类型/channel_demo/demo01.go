package main

import "fmt"

type worker interface {
	work()
}

type person struct {
	worker
}

func Reverse(data worker) worker {
	return &person{data}
}

func (p person) work() {
	fmt.Println("work....")
}

func main() {
	var w worker = person{}
	Reverse(w).work()
}
