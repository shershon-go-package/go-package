package test

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"net/http"
	"shershon1991/go-study-example/app/collyDemo"
	"testing"
)

func TestCollyDemo(t *testing.T) {
	err := collyDemo.RunDemo()
	if err != nil {
		t.Error(err)
	}
}

func TestDouBan(t *testing.T) {
	err := collyDemo.DouBanBook()
	if err != nil {
		t.Error(err)
	}
}

func TestJiJIn(t *testing.T) {
	url := "https://www.morningstar.cn/quickrank/rqfii.aspx"
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"),
		colly.Debugger(&debug.LogDebugger{}),
		//colly.Async(true),
	)
	cookieList := []*http.Cookie{
		&http.Cookie{
			Name:     "authWeb",
			Value:    "AFE9FB65A838E022E595BC294AD98AE8B78C2DE6FE4484976A8E032513B0E4FAEC101AF2D212BE94AD36A029DAAF436615769E72A0966B8A97C556FB741ADAC6C4F706B0C427DFBB53F7A6E176E7E4D21E525BE38877C3657721969286FFCA34168496E2087E8C11692A7D0DEABD72818CAD8038",
			Path:     "/",
			Domain:   ".morningstar.cn",
			Secure:   true,
			HttpOnly: true,
		},
	}
	err := collector.SetCookies(url, cookieList)
	if err != nil {
		t.Errorf("set cookie error:%s", err)
		return
	}
	collector.OnError(func(response *colly.Response, err error) {
		t.Errorf("req error:%s", err)
		return
	})
	collector.OnHTML("#ctl00_cphMain_gridResult", func(element *colly.HTMLElement) {
		fmt.Println("tbody:", element.ChildText("tbody"))
	})
	err = collector.Visit(url)
	collector.OnError(func(response *colly.Response, err error) {
		t.Errorf(" error:%s", err)
		return
	})
	//collector.Wait()
}

// 返回当前元素的属性
func TestUseAttr(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("div[class='nav-logo'] > a", func(element *colly.HTMLElement) {
		// 定位到div[class='nav-logo'] > a标签元素
		fmt.Printf("href:%v\n", element.Attr("href"))
	})
	_ = collector.Visit("https://book.douban.com/tag/小说")
}

// 测试使用ChildAttr和ChildAttrs
func TestChildAttrMethod(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("OnError", err)
	})
	//
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// 获取第一个子元素(div)的class属性
		fmt.Printf("ChildAttr:%v\n", element.ChildAttr("div", "class"))
		// 获取所有子元素(div)的class属性
		fmt.Printf("ChildAttrs:%v\n", element.ChildAttrs("div", "class"))
	})
	err := collector.Visit("https://liuqh.icu/a.html")
	if err != nil {
		fmt.Println("err", err)
	}
}

// 测试使用ChildText和ChildTexts
func TestChildTextMethod(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("OnError", err)
	})
	//
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// 获取第一个子元素(div)的class属性
		fmt.Printf("ChildText:%v\n", element.ChildText("div"))
		// 获取所有子元素(div)的class属性
		fmt.Printf("ChildTexts:%v\n", element.ChildTexts("div"))
	})
	err := collector.Visit("https://liuqh.icu/a.html")
	if err != nil {
		fmt.Println("err", err)
	}
}

// 遍历
func TestForeach(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("ul[class='demo']", func(element *colly.HTMLElement) {
		element.ForEach("li", func(_ int, el *colly.HTMLElement) {
			name := el.ChildText("span[class='name']")
			age := el.ChildText("span[class='age']")
			home := el.ChildText("span[class='home']")
			fmt.Printf("姓名: %s  年龄:%s 住址: %s \n", name, age, home)
		})
	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}

// 定义结构体
type Book struct {
	Name    string   `selector:"span.title"`
	Link    string   `selector:"span > a" attr:"href"`
	Author  string   `selector:"span.autor"`
	Reviews []string `selector:"ul.category > li"`
	Price   string   `selector:"span.price"`
}

func TestUnmarshal(t *testing.T) {
	// 声明结构体
	var book Book
	collector := colly.NewCollector()
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		err := element.Unmarshal(&book)
		if err != nil {
			fmt.Println("解析失败:", err)
		}
		fmt.Printf("结果:%+v\n", book)
	})
	fmt.Printf("book:%+v\n", book)
	_ = collector.Visit("https://liuqh.icu/a.html")
}

// 测试提取html.table
func TestCollectTable(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("div[class='list'] table", func(element *colly.HTMLElement) {
		element.ForEach("tr", func(i int, el *colly.HTMLElement) {
			// 过滤表头
			if i == 0 {
				return
			}
			name := el.ChildText("td:nth-of-type(1)")
			age := el.ChildText("td:nth-of-type(2)")
			home := el.ChildText("td:nth-of-type(3)")
			fmt.Printf("姓名:%s 年龄:%s 地址:%s \n", name, age, home)
		})
	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}

// 常用选择器使用
func TestSelector(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// ID属性选择器,使用#
		fmt.Printf("ID选择器使用: %v \n", element.ChildText("#title"))
		// Class属性选择器,使用
		fmt.Printf("class选择器使用1: %v \n", element.ChildText("div[class='desc']"))
		fmt.Printf("class选择器使用2: %v \n", element.ChildText(".desc"))
		// 相邻选择器 prev +  next: 提取 <span>好好学习！</span>
		fmt.Printf("相邻选择器: %v \n", element.ChildText("div[class='desc'] + span"))
		// 父子选择器： parent > child,提取:<div class="parent">下所有子元素
		fmt.Printf("父子选择器: %v \n", element.ChildText("div[class='parent'] > p"))
		// 兄弟选择器 prev ~ next , 提取:<p class="childB">老二</p>
		fmt.Printf("兄弟选择器: %v \n", element.ChildText("p[class='childA'] ~ p"))
		// 同时选中多个,用,
		fmt.Printf("同时选中多个1: %v \n", element.ChildText("span[class='context1'],span[class='context2']"))
		fmt.Printf("同时选中多个2: %v \n", element.ChildText(".context1,.context2"))
	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}

// 常用过滤器使用first-child & first-of-type
func TestFilterFirstChild(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// 只会筛选父元素下第一个子元素是<p>..</p>
		fmt.Printf("first-child: %v \n", element.ChildText("p:first-child"))
		// 会筛选父元素下第一个子元素类型是<p>..</p>
		fmt.Printf("first-of-type: %v \n", element.ChildText("p:first-of-type"))
	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}

// 过滤器第x个元素
func TestFilterNth(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		//<div class="parent">下的第一个子元素
		nthChild := element.ChildText("div[class='parent'] > :nth-child(1)")
		fmt.Printf("nth-child(1): %v \n", nthChild)

		//<div class="parent">下的第一个p子元素
		nthOfType := element.ChildText("div[class='parent'] > p:nth-of-type(1)")
		fmt.Printf("nth-of-type(1): %v \n", nthOfType)

		// div class="parent">下的最后一个子元素
		nthLastChild := element.ChildText("div[class='parent'] > :nth-last-child(1)")
		fmt.Printf("nth-last-child(1): %v \n", nthLastChild)

		//<div class="parent">下的最后一个p子元素
		nthLastOfType := element.ChildText("div[class='parent'] > p:nth-last-of-type(1)")
		fmt.Printf("nth-last-of-type(1): %v \n", nthLastOfType)

	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}

// 过滤器只有一个元素
func TestFilterOnly(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// 匹配其子元素：有且只有一个标签的
		onlyChild := element.ChildTexts("div > :only-child")
		fmt.Printf("onlyChild: %v \n", onlyChild)
		// 匹配其子元素：有且只有一个 p 标签的
		nthOfType := element.ChildTexts("div > p:only-of-type")
		fmt.Printf("nth-of-type(1): %v \n", nthOfType)

	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}

func TestFilterContext(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// 内容匹配
		attr1 := element.ChildAttr("a:contains(百度)", "href")
		attr2 := element.ChildAttr("a:contains(必应)", "href")
		fmt.Printf("百度: %v \n", attr1)
		fmt.Printf("必应: %v \n", attr2)
	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}
