package 第五章数据

import (
	"bytes"
	"strings"
)

func StringJoin() string {
	s := make([]string, 1000)
	for i := 0; i < len(s); i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}

func StringJoin02() string {
	var b bytes.Buffer
	b.Grow(10)
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}
