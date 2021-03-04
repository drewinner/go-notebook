package sort

import (
	"fmt"
	"testing"
)

func TestSelectMax(t *testing.T) {
	max := SelectMax([]int{9, 3, 1, 2, 4})
	fmt.Println(max)
	//选择排序
	arr := SelectSort([]int{9, 3, 1, 2, 4})
	fmt.Println(arr)
}
