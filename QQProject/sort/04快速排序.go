package sort

import (
	"fmt"
)

/**
*	快速排序
 */
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]

	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)
	mid = append(mid, pivot)

	for i := 1; i < len(arr); i++ {
		if arr[i] > pivot { //当前元素大于轴点、放到右边
			high = append(high, arr[i])
			continue
		}
		if arr[i] < pivot {
			low = append(low, arr[i])
			continue
		}
		mid = append(mid, arr[i])
	}
	low, high = QuickSort(low), QuickSort(high)
	fmt.Println(low, high, mid)
	myArr := append(append(low, mid...), high...)
	return myArr
}

func QuickSortV2(arr []int, begin, end int) []int {

	if end-begin < 2 {
		return arr
	}
	//确定轴点位置
	mid := pivotIndex(arr, begin, end)
	QuickSortV2(arr, begin, mid)
	QuickSortV2(arr, mid+1, end)
	return arr
}

func pivotIndex(arr []int, begin, end int) int {
	//备份轴点元素的位置
	pivot := arr[begin]
	//end指向最后一个元素
	end--
	for begin < end {
		//先从右边元素和轴点元素进行对比、
		//1. 如果右边元素大于轴点元素end--、end和begin交换位置、begin++
		for begin < end {
			if arr[end] > pivot {
				end--
			} else {
				//右边元素覆盖左边元素
				arr[begin] = arr[end]
				begin++
				break
			}
		}
		for begin < end {
			//2.从左边开始
			//3. 如果左边元素小于轴点元素
			if arr[begin] < pivot {
				begin++
			} else {
				//左边元素覆盖右边元素
				arr[end] = arr[begin]
				end--
				break
			}
		}

	}
	//4.将轴点元素放入最终的位置
	arr[begin] = pivot
	//返回轴点元素的位置
	return begin
}
