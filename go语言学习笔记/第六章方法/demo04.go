package 第六章方法

import (
	"fmt"
	"reflect"
)

type SS struct {
}

type TT struct {
	SS
}

func (SS) sVal()  {}
func (*SS) sPtr() {}
func (TT) tVal()  {}
func (*TT) tPtr() {}

//显示方法集里所有方法名称
func MethodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("==", t.NumMethod(), "==")
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
