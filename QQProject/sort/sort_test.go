package sort

import (
	"fmt"
	"testing"
)

func TestSelectMax(t *testing.T) {
	//max := SelectMax([]int{9, 3, 1, 2, 4})
	//fmt.Println(max)
	////选择排序
	//arr := SelectSort([]int{9, 3, 1, 2, 4})
	//fmt.Println(arr)
	//test := QuickSort([]int{1, 23, 2, 44, 22, 11})
	arr := []int{1, 23, 2, 44, 22, 11}
	test := QuickSortV2(arr, 0, len(arr))
	fmt.Println(test)
}
