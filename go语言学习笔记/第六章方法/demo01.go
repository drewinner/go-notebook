package 第六章方法

import "fmt"

type N int

func (n N) Chap6Demo01() string {
	return fmt.Sprintf("%#x", n)
}
