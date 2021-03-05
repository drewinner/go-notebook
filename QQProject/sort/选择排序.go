package sort

func SelectMax(arr []int) (max int) {
	if len(arr) <= 1 {
		return arr[0]
	}
	//选择最大的值
	max = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i] //始终存储最大的
		}
	}
	return max
}

/**
*	选择排序
 */
func SelectSort(arr []int) (sortArr []int) {
	if len(arr) <= 1 {
		return arr
	}
	//外层控制选择的位置
	for i := 0; i < len(arr)-1; i++ {
		min := i //存放最小数的索引
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min] //找到最小的、位置进行交换
		}
	}
	return arr
}
