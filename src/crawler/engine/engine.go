package engine

import (
	"fmt"
	"log"

	"hellosz.top/src/crawler/fetcher"
	"hellosz.top/src/crawler/types"
)

// Run 爬虫项目的入口，输入种子数据，然后使用深度遍历算法进行爬取
func Run(seeds []types.Request) {
	// 初始化种子数据
	var requests []types.Request
	requests = append(requests, seeds...)

	// 遍历商品信息
	cities := 0
	for len(requests) > 0 {
		// 取出队首
		request := requests[0]
		requests = requests[1:]

		// 调用fetcher，进行数据爬取
		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fetch url(%s) failed, reason: %s", request.Url, err)
		}

		// 调用parser，进行数据解析
		parseResult := request.ParserFunc(body)

		// 返回结果进行入队
		requests = append(requests, parseResult.Requests...)
		cities++
	}

	// 输出结果
	fmt.Printf("一共爬取了%d个城市", cities)
}
