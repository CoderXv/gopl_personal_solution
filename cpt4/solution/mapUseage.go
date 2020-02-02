package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// 4.8 修改charcount的代码来统计字母,数字和其他Unicode分类中的字符数量,可以使用函数unicode.IsLetter等
// EOF: return then control + d in macOS, control + z in windows
func Charcount() {
	counts := make(map[string]int)    // counts of Unicode characters
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			counts["letter"] += 1
			continue
		}
		if unicode.IsNumber(r) {
			counts["number"] += 1
			continue
		}
		if unicode.IsControl(r) {
			counts["control"] += 1
			continue
		}
		counts["other"] += 1
	}
	fmt.Printf("type\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// 4.9 编写一个程序wordfreq来汇总输入文本文件中每个单词出现的次数.在第一次调用scan之前,
// 需要使用input.Split(bufio.ScanWords)来将文本行按照单词分割而不是行分割
// 参考文献 https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4
func Wordfreq() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	// 需要使用input.Split(bufio.ScanWords)来将文本行按照单词分割而不是行分割
	scanner.Split(bufio.ScanWords)
	record := make(map[string]int)
	for scanner.Scan() {
		record[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input:", err)
	}
	fmt.Printf("%s\t%s\n", "Word", "Count")
	for k, v := range record {
		fmt.Printf("%s\t%d\n", k, v)
	}
}