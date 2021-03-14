package 常用题目

/**
*	https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
*	找出数组中重复的数字。
*	在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
*	也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
*	输入：
*	[2, 3, 1, 0, 2, 5, 3]
*	输出：2 或 3
 */
func FindRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		} else {
			m[num] = true
		}
	}
	return 0
}

func FindRepeatNumber01(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//控制扫描位置
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			//交换
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}
