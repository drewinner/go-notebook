package 常用题目

/**
*	https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
*	输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
*	输入: n = 1
*	输出: [1,2,3,4,5,6,7,8,9]
 */
func printNumbers(n int) []int {
	sum := 1
	for i := 0; i < n; i++ {
		sum = sum * 10
	}
	sum--
	s := make([]int, sum)
	for i := 0; i < sum; i++ {
		s[i] = i + 1
	}
	return s
}
