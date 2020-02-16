package solution
import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"unicode"
)


// 4.1 编写一个函数,计算两个SHA256哈希中不同的bit的数目
func Sha256Count(s1 []byte, s2[]byte) int {
	if len(s1) != len(s2) {
		return 0
	}

	var res = 0

	for i := range s1 {
		if s1[i] == s2[i] {
			res++
		}
		i++
	}

	return res
}

// 4.2 编写一个程序,默认情况下打印标准输入的SHA256编码,并支持通过命令行flag定制,输出SHA384或SHA512哈希算法
func ShaOutput() {
	method := flag.String("method", "sha256", "select hash method(sha256,sha384,sha512)")
	text := flag.String("text", "", "input text string you want to hash")

	// 使用parse初始化函数
	flag.Parse()

	switch *method {
	case "sha256":
		fmt.Printf("Sha256: %x\n", sha256.Sum256([]byte(*text)))
	case "sha384":
		fmt.Printf("Sha384: %x\n", sha512.Sum384([]byte(*text)))
	case "sha512":
		fmt.Printf("Sha512: %x\n", sha512.Sum512([]byte(*text)))
	default:
		fmt.Println("Error Input, Please Check.")
	}
}

// 4.3 重写函数reverse,使用数组指针作为参数而不是slice
func ReverseRemaster(s *[5]int) {
	for i, j := 0, len(*s) - 1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

// 4.4 编写一个函数rorate,实现一次遍历就可以完成元素旋转
// 这里需要注意的是输入s并没有修改，因为生成新的slice并返回。
func RotateOnce(s *[]int, index int) []int {
	r := (*s)[index:]
	for i := index - 1; i >= 0; i -- {
		r = append(r, (*s)[i])
	}
	return r
}

// 4.5 编写一个就地处理函数,用于去除[]string slice中相邻的重复字符串元素
func RmSameChar(s []string) []string {
	if len(s) < 1 {
		return s
	}
	current := 0 // 当前下标
	position := 0 // 处理完成的下标
	for ; current < len(s); current++ {
		if s[current] != s[position] {
			position += 1
			s[position] = s[current]
		}
	}
	return s[:position + 1]
}

// 4.6 编写一个就地处理函数,用于将一个UTF-8编码的字节slice中所有相邻的Unicode空白字符(查看unicode.IsSpace)缩减为一个ASCII空白字符
func RmUniSame(s []rune) []rune {
	position := 0
	tag := unicode.IsSpace(s[0])
	for i := 1; i < len(s); i++ {
		if unicode.IsSpace(s[i]) {
			if tag != true {
				position += 1
				s[position] = s[i]
				tag = true
			}
		} else {
			tag = false
			position += 1
			s[position] = s[i]
		}
	}
	return s[:position+1]
}

// 4.7 修改函数reverse,来翻转一个UTF-8编码的字符串中的字符元素,传入参数是该字符串对应的字节slice类型([]byte).
// 你可以做到不需要重新分配内存就实现该功能吗 (暂时做不到)
/*
	这个感觉有点难,类型转换会导致内存重新分配
	我的思路是对每个byte读取前面有几个1，有几个1则表示有多少个byte是连续的。
至于翻转，感觉不好翻…..但是能够转换类型就好办了，直接转换成rune
*/
func RevRemaster(s []byte) []byte {
	b := []rune(string(s))
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return []byte(string(b))
}
