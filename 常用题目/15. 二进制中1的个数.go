package 常用题目

/**
*	https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
*	请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。
*	例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
*	输入：00000000000000000000000000001011
*	输出：3
*	解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
 */
func hammingWeight(num uint32) int {
	var rs uint32 = 0
	for num > 0 {
		rs += num & 1
		num = num >> 1
	}
	return int(rs)
}
