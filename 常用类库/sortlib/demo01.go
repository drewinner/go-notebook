package sortlib

import (
	"fmt"
	"sort"
)

func SortLibDemo01() {
	data := []int{10, 25, 11, 24, 17, 26}
	i := sort.Search(len(data), func(i int) bool {
		return data[i] >= 23
	})
	fmt.Println("最终的结果为", i) //
}

func SortLibDemo02() {
	es := []string{"abc"}
	s := "abc/def"
	i := sort.Search(len(es), func(i int) bool {
		return len(es[i]) < len(s)
	})
	fmt.Println(i)
}
