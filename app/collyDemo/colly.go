package collyDemo

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

func RunDemo() error {
	// 创建 Collector 对象
	collector := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"))
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("在请求之前调用")
	})
	// 请求期间发生错误,则调用
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("请求错误:", err)
	})
	// 收到响应后调用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("url:", response.Request.URL)
	})
	//OnResponse如果收到的内容是HTML ,则在之后调用
	collector.OnHTML("#position_shares table", func(element *colly.HTMLElement) {
		// 遍历table
		element.ForEach("table tr", func(_ int, el *colly.HTMLElement) {
			name := el.ChildText("td:nth-of-type(1)")
			percentage := el.ChildText("td:nth-of-type(2)")
			fmt.Printf("名称:%v 仓位占比:%v \n", name, percentage)
		})

	})
	return collector.Visit("https://fund.eastmoney.com/481010.html")
}

// 豆瓣书榜单
func DouBanBook() error {
	// 创建 Collector 对象
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"),
		colly.CacheDir("/tmp/"),
	)
	//collector := colly.NewCollector()
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("回调函数OnRequest: 在请求之前调用")
	})
	// 请求期间发生错误,则调用
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("回调函数OnError: 请求错误", err)
	})
	// 收到响应后调用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("回调函数OnResponse: 收到响应后调用")
	})
	//OnResponse如果收到的内容是HTML ,则在之后调用
	collector.OnHTML("ul[class='subject-list']", func(element *colly.HTMLElement) {
		// 遍历li
		element.ForEach("li", func(i int, el *colly.HTMLElement) {
			coverImg := el.ChildAttr("div[class='pic'] > a[class='nbg'] > img", "src")
			bookName := el.ChildText("div[class='info'] > h2")
			authorInfo := el.ChildText("div[class='info'] > div[class='pub']")
			split := strings.Split(authorInfo, "/")
			author := split[0]
			fmt.Printf("封面: %v 书名:%v 作者:%v\n", coverImg, trimSpace(bookName), author)
		})

	})
	return collector.Visit("https://book.douban.com/tag/小说")
}

func trimSpace(str string) string {
	// 替换所有的空格
	str = strings.ReplaceAll(str, " ", "")
	// 替换所有的换行
	return strings.ReplaceAll(str, "\n", "")
}
