package 第五章数据

import (
	"fmt"
	"unsafe"
)

func Demo06() {
	var a struct{}
	var b [100]struct{}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}
