package 常用题目

/**
*	https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
*	一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
*	答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
 */
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	first := 0
	second := 1

	for i := 0; i < n; i++ {
		sum := first + second
		first = second % 1000000007
		second = sum % 1000000007
	}
	return second
}
