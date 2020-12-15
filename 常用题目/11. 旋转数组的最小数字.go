package 常用题目

/**
*	https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
*	把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
*	输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
*	输入：[3,4,5,1,2]
*	输出：1
*	输入：[2,2,2,0,1]
*	输出：0
*	输入：[1,3,5]
*	输出：1
 */
func MinArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	flag := 0 //0表示一直递增、1表示有递减
	rs := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] <= numbers[i+1] {
			continue
		} else {
			flag = 1
		}
		rs = numbers[i+1]
	}
	if flag == 0 {
		return numbers[0]
	} else {
		return rs
	}
}
