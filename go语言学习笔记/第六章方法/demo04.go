package 第六章方法

import (
	"fmt"
	"reflect"
)

type SSInterface interface {
	sVal()
	sPtr()
}

type SS struct {
}

func (s *SS) sVal() {}
func (s *SS) sPtr() {}

//显示方法集里所有方法名称
func MethodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("==", t.String(), t.NumMethod(), "==")
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
