package 第五章数据

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PP(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}
