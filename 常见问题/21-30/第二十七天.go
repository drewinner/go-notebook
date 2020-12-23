package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type Math struct {
	x, y int
}

var m = map[string]*Math{
	"foo": &Math{2, 3},
}

//go 中的 map 的 value 本身是不可寻址的
func main() {
	fmt.Println(South)
	m["foo"].x = 100
	//tmp := m["foo"]
	//tmp.x = 4
	//m["foo"] = tmp
	fmt.Println(m["foo"].x)
	i := "world"
	m := make(map[int]*string)
	m[0] = &i
	fmt.Println(m[0])
}
