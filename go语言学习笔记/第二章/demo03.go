package 第二章

func Makeslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func MakeMap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}
