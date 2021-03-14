package 常用题目

import (
	"strings"
)

/**
*	https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
*	请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*	输入：s = "We are happy."
*	输出："We%20are%20happy."
 */
//原有字符串基础上修改
func ReplaceSpace03(s string) string {
	return strings.Replace(s, " ", "%20", -1)

}

func ReplaceSpace(s string) string {
	rs := ""
	for _, v := range s {
		if string(v) == " " {
			rs = rs + "%20"
			continue
		}
		rs = rs + string(v)

	}
	return rs
}
func ReplaceSpace02(s string) string {
	if s == "" {
		return s
	}
	//计算空格数量
	spaceCount := 0
	for _, v := range s {
		if string(v) == " " {
			spaceCount++
		}
	}
	if spaceCount == 0 {
		return s
	}
	//初始化字节数组
	b := make([]byte, len(s)+spaceCount*2)
	size := 0
	for _, v := range s {
		if string(v) == " " {
			b[size] = byte('%')
			size++
			b[size] = byte('2')
			size++
			b[size] = byte('0')

		} else {
			b[size] = byte(v)
		}
		size++
	}
	return string(b)
}
