package 第六章方法

import (
	"fmt"
	"testing"
)

func TestN_Chap6Demo01(t *testing.T) {
	//var x *X
	//x.Test()
	//X{}.Test()
	var s SS
	MethodSet(s)
	fmt.Println("++++")
	MethodSet(&s)
	//u := user{"Bill", "bill@email.com"}
	//sendNotificatioin(u)
}
