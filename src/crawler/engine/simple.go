package engine

import (
	"hellosz.top/src/crawler/fetcher"
	"hellosz.top/src/crawler/types"
	"log"
)

// Simple 单进程爬虫
type Simple struct{}

// Run 爬虫项目的入口，输入种子数据，然后使用深度遍历算法进行爬取
func (*Simple) Run(seeds []types.Request) {
	// 初始化种子数据
	var requests []types.Request
	requests = append(requests, seeds...)

	// 遍历商品信息
	for len(requests) > 0 {
		// 取出队首
		request := requests[0]
		requests = requests[1:]

		// 爬取请求
		parseResult, err := Worker(request)
		if err != nil {
			panic(err)
		}

		// 返回结果进行入队
		requests = append(requests, parseResult.Requests...)
	}
}

// Worker 爬虫工作者
func Worker(request types.Request) (types.ParseResult, error) {
	// 调用fetcher，进行数据爬取
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch url(%s) failed, reason: %s", request.Url, err)

		return types.ParseResult{}, err
	}

	// 调用parser，进行数据解析
	parseResult := request.ParserFunc(body)
	return parseResult, nil
}
