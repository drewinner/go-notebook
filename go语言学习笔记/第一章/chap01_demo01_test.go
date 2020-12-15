package 第一章

import (
	"fmt"
	"testing"
)

func TestUser_Print(t *testing.T) {
	var m manager
	m.Name = "Tom"
	m.Age = 11
	fmt.Println(m.ToString())
}

func TestUser02_Print(t *testing.T) {
	var u user02
	u.name = "Tom"
	u.age = 111
	var p Printer = u
	p.Print()
}
