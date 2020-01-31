package solution

import (
	"bytes"
	"fmt"
	"strings"
)

// 3.10 编写一个非递归的comma函数，运用bytes.Buffer, 而不是简单的字符拼接
// comma 从右侧开始每三位数字后就插入一个逗号，仅考虑整数
func Comma(s string) string {
	var buffers bytes.Buffer

	n := len(s)
	if n <= 3 {
		return s
	}

	mod := n % 3
	if mod > 0 {
		buffers.Write([]byte(s[:mod] + ","))
	}
	for mod + 3 < n {
		buffers.Write([]byte(s[mod:mod+3] + ","))
		mod += 3
	}
	if mod + 3 == n {
		buffers.Write([]byte(s[mod:mod+3]))
	}

	return buffers.String()
}

// 3.11 增强comma函数的功能，让其正确处理浮点数，以及带有可选正负号的数字
func CommaFloatFlag(s string) string {

	flag := string(s[0])
	splitRes := strings.Split(s, ".")

	prefix := splitRes[0][1:]
	postfix := splitRes[1]

	prefixStr := Comma(prefix)

	return strings.Join([]string{flag, prefixStr, postfix}, "")
}


// 3.12 编写一个函数判断两个字符串是否同文异构，也就是，它们都含有相同的字符但排列顺序不同
func Isomerism (s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	// 遍历s1，统计每个字符在整个数组中出现的频率是否与s2一样
	for _, v := range s1 {
		// 注意如果v的类型为int32，需要用string()转为字符
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}

	return true
}

// 3.13 用尽可能简洁的方法申明从KB MB 直到YB的常量
const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)