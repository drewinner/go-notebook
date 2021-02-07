package 常用题目

/**
*	https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
*	0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
*	例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
 */

//func lastRemaining(n int, m int) int {
//
//	l := list.New()
//
//	for i := 0; i < n; i++ {
//		l.PushBack(i)
//	}
//	current := l.Front() //获取第一个元素
//	next := current.Next()
//	delIndex := 1
//	for l.Len() > 1 {
//		if delIndex%m == 0 {
//			next = current.Next()
//			l.Remove(current)
//		}
//		delIndex++
//		if next == nil {
//			current = l.Front()
//		} else {
//			current = next
//		}
//		next = current.Next()
//	}
//	return current.Value.(int)
//}

//func lastRemaining(n int, m int) int {
//	delIndex := 0
//
//	numSlice := make([]int, n)
//
//	for i := 0; i < n; i++ {
//		numSlice[i] = i
//	}
//	for len(numSlice) > 1 {
//		delIndex = delIndex + m - 1
//		if delIndex >= len(numSlice) {
//			delIndex = delIndex % len(numSlice)
//		}
//		numSlice = append(numSlice[0:delIndex], numSlice[delIndex+1:]...)
//	}
//
//	return numSlice[0]
//}
func lastRemaining(n int, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
