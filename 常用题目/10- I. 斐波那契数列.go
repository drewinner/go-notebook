package 常用题目

/**
*	https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
*	写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下
*	F(0) = 0,   F(1) = 1
*	F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
*	斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
*	答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
 */
func Fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	first := 0
	second := 1
	for i := 0; i < n-1; i++ {
		sum := first + second
		first = second % 1000000007
		second = sum % 1000000007
	}
	return second
}

func Fib2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib2(n-1) + Fib2(n-2)
}
