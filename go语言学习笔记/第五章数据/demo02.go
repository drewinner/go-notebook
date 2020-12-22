package 第五章数据

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Demo02() {
	s := "abcdefg"
	s1 := s[:3]
	s2 := s[1:4]
	s3 := s[2:]
	fmt.Println(s1, s2, s3)
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
	var b []byte
	b = append(b, "abc"...)
	fmt.Println(b)
}
