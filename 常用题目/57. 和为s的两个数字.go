package 常用题目

/**
*	https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
*	输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
*	输入：nums = [2,7,11,15], target = 9
*	输出：[2,7] 或者 [7,2]
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) //值->索引
	for _, v := range nums {
		index := target - v
		if vv, ok := m[index]; ok {
			return []int{vv, v}
		}
		m[v] = v
	}
	return []int{}
}
