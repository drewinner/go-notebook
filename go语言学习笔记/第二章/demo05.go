package 第二章

import "fmt"

type Dint int

var aaa = 10

func Test05() {
	var d Dint = 10
	dd := int(d)
	fmt.Println(dd)
	//var x int = d
	//fmt.Println(d == x)
	x := 1
	switch x {
	case 1:
		fallthrough
	case 2:
		fmt.Println(2)
	}
	fmt.Println(&aaa)
	_, aaa := 20, 2
	fmt.Println(&aaa)
}
