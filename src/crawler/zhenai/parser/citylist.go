package parser

import (
	"bytes"
	"crawler/types"
	"github.com/PuerkitoBio/goquery"
	"log"
)

// ParseCityList 解析城市列表
func ParseCityList(contents []byte) types.ParseResult {
	// 读取其中的节点
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		panic(err)
	}

	// 解析结果
	var result types.ParseResult
	dom.Find(".city-list>dd>a").Each(func(i int, selection *goquery.Selection) {
		url := selection.AttrOr("href", "nil href")
		text := selection.Text()

		// 缓存请求结果
		result.Items = append(result.Items, text)
		result.Requests = append(result.Requests, types.Request{
			Url:        url,
			ParserFunc: types.NilParser,
		})

		// 打印解析结果
		log.Printf("saving city:%s\n", text)
	})

	return result
}
