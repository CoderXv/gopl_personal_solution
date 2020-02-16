package solution

import (
"fmt"
"log"
"os"
	"sort"
	"time"
	"tutorial/cpt4/github"
)

// 4.10 修改issues实例,按照时间来输出结果,比如一个月以内,一年以内或者超过一年
func Issues() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

// 对创建日期进行排序
func SortResult(result *github.IssuesSearchResult) {
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.Before(result.Items[j].CreatedAt)
	})
}

// 4.10 main process
func SearchIssuesMainProcess() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	SortResult(result)
	var month, year, distant []*github.Issue

	// 基于当前日期计算时间范围
	for v, i := range result.Items {
		if i.CreatedAt.After(time.Now().Add(time.Hour * 24 * 30 * -1)) {
			month = append(month, result.Items[v])
		} else if i.CreatedAt.After(time.Now().Add(time.Hour * 24 * 365 * -1)) {
			year = append(year, result.Items[v])
		} else {
			distant = append(distant, result.Items[v])
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("with in a month\n")
	for _, item := range month {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("with in a year\n")
	for _, item := range year {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Printf("long long ago\n")
	for _, item := range distant {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

// 4.11 开发一个工具让用户可以通过命令行创建、读取、更新或者关闭Github的issues,当需要额外输入的时候,调用他们喜欢的文本编辑器
// omit

// 4.12 流行web漫画xkcd有一个JSON接口.例如,调用https://xkcd.com/571/info.0.json输出漫画571的详细描述,这个是很多人最喜欢的之一.
// 下载每一个URL并且构建一个离线索引.编写xkcd来使用这个索引,可以通过命令行指定的搜索条件来查找并输出符合条件的每个漫画的URL和剧本

// 4.13 基于JSON开发的web服务,开放电影数据库让你可以在https://omdbapi.com/上通过名字来搜索电影并下载海报图片.
// 开发一个poster工具以通过命令行指定的电影名称来下载海报
// omit


