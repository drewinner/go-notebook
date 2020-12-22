package 第六章方法

import (
	"fmt"
	"testing"
)

func TestN_Chap6Demo01(t *testing.T) {
	//var x *X
	//x.Test()
	//X{}.Test()
	var tt TT
	MethodSet(tt)
	fmt.Println("++++")
	MethodSet(&tt)
}
