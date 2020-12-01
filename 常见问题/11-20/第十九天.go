package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age) //28

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person) //29

	// 3.
	defer func() {
		fmt.Println(person.age) //29
	}()

	person.age = 29
}
