// Package provides ...
package parser

import (
	"bytes"
	// "fmt"
	"github.com/PuerkitoBio/goquery"
	"hellosz.top/src/crawler/types"
	"log"
)

const userSelector = ".list-item > .content a"
const nextPageSelector = ".f-pager li a"
const nextPageTag = "下一页"

// ParseCity 解析城市
func ParseCity(contents []byte) types.ParseResult {
	// 读取其中的节点
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		panic(err)
	}

	// 解析当前城市页面的用户链接
	var result types.ParseResult
	dom.Find(userSelector).Each(func(i int, selection *goquery.Selection) {
		url := selection.AttrOr("href", "nil href")
		text := selection.Text()

		// 缓存请求结果
		result.Items = append(result.Items, text)
		result.Requests = append(result.Requests, types.Request{
			Url:        url,
			ParserFunc: ParseProfile,
		})

		// 打印解析结果
		log.Printf("saving user:%s\n", text)
	})

	// // 解析下一页
	// pageSelection := dom.Find(nextPageSelector).Last()
	// fmt.Println(pageSelection.Html())
	// url := pageSelection.AttrOr("href", "nil href")
	// text := pageSelection.Text()
	// if text == nextPageTag {
	// 	// 缓存请求结果
	// 	result.Items = append(result.Items, text)
	// 	result.Requests = append(result.Requests, types.Request{
	// 		Url:        url,
	// 		ParserFunc: ParseCity,
	// 	})
	// 	log.Printf("saving city next page:%s\n", url)
	// }

	return result
}
