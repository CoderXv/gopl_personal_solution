package solution

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 5.1 改变findlinks程序,使用递归调用visit(而不是循环)遍历n.FirstChild链 (不会)
func Findlinks() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// 5.2 写一个函数,用于统计HTML文档树内所有的元素个数,如p,div,span等
func visitElem(e map[string]int, n *html.Node) []string {
	if n.Type == html.ElementNode {
		e[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitElem(e, c)
	}
}

// 5.3 写一个函数,用于输出HTML文档树种所有文本节点的内容.但不包括<script>或<style>元素,因为这些内容在web浏览器中是不可见的
func visitTextContent(n *html.Node) {
	if n != nil && n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitTextContent(c)
	}
}

// 5.4 扩展visit函数,使之能够获得到其他种类的链接地址,比如图片、脚本或样式表的链接
var filter = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
}

func visitManyElem(links []string, n *html.Node) []string {
	for k, v := range filter {
		if n.Type == html.ElementNode && n.Data == k {
			for _, a := range n.Attr {
				if a.Key == v {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitManyElem(links, c)
	}
	return links
}

// 5.5 实现函数countWordsAndImages
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode{
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ws, is := countWordsAndImages(c)
		words += ws
		images += is
	}
	return words, images
}

// 5.6 修改gopl.io/ch3/surface中的函数corner,以使用命名的结果以及裸返回语句
func corner(i, j int) (sx, sy float64) {
	return
}

// 5.7 开发startElement和endElement函数并应用到一个普通的HTML输出代码中.输出注释节点、文本节点和所有元素属性(<a href='...'>).
// 当一个元素没有子节点时,使用简短的形式,比如<img/>而不是<img></img>.写一个测试程序保证输出可以正确解析

// 5.8 修改forEachNode使得pre和post函数返回一个布尔型的结果来确定遍历是否继续下去.使用它写一个函数ElementByID,
// 该函数使用下面的函数签名并且找到第一个符合id属性的HTML元素.函数在找到符合条件的元素时应该尽快停止遍历.func
// 使用它写一个函数ElementByID(doc *html.Node, id string) *html.Node

// 5.9 写一个函数expand(s string,f func(string)string)string,该函数替换参数s中每一个子字符串”$foo”为f(“foo”)的返回值

// 5.10 重写topSort以使用map代替slice并去掉开头的排序.结果不是唯一的,验证这个结果是合法的拓扑排序

// 5.11 现在有”线性代数”这门课程,它的先决课程是”微积分”.拓展topSort以函数输出结果

// 5.12 5.5节(gopl.io/ch5/outline2)的startElement和endElement函数共享一个全局变量depth.把它们变为匿名函数以共享outline函数的一个局部变量

// 5.13 修改crawl函数保存找到的页面,根据需要创建目录.不要保存不同域名下的页面.比如本来的页面来自golang.org,那么就把它们保存下来但是不要保存vimeo.com下的页面

// 5.14 使用广度优先遍历搜索一个不同的拓扑结构.比如,你可以借鉴拓扑排序的例子里的课程依赖关系,计算机文件系统的分层结构,或者从当前城市的官网上下载公共汽车或者地铁的路线图.

// 5.15 模仿sum写两个变长函数max和min.当不带任何参数调用这些函数的时候应该怎么应对?编写类似函数的变种,要求至少需要一个参数

// 5.16 写一个变长版本的strings.Join函数

// 5.17 写一个变长函数ElementsByTagname,已知一个HTML节点树和零个或多个名字,返回所有符合给出名字的元素

// 5.18 不改变原本的行为,重写fetch函数以使用defer语句关闭打开的可写的文件

// 5.19 使用panic和recover写一个函数,它没有return语句,但是能够返回一个非零的值

