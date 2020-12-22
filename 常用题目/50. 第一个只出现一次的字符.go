package 常用题目

/**
*	https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
*	在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
*	s = "abaccdeff"
*	返回 "b"
*	s = ""
*	返回 " "
 */
func FirstUniqChar01(s string) byte {
	if s == "" {
		return ' '
	}
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		v := s[i]
		m[v]++
	}
	//二次遍历
	for i := 0; i < len(s); i++ {
		v := s[i]
		if m[v] == 1 {
			return v
		}
	}

	return ' '
}

func FirstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		v := s[i]
		m[v]++
	}
	//二次遍历
	for i := 0; i < len(s); i++ {
		v := s[i]
		if m[v] == 1 {
			return v
		}
	}

	return ' '
}
