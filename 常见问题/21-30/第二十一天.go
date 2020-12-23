package main

type S struct {
}

func f(x interface{}) {
}

//永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。
func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	//g(s) //B
	f(p) //C
	//g(p) //D
}
