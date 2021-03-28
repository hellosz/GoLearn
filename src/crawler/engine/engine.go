package engine

import (
	"crawler/fetcher"
	"crawler/types"
	"crawler/zhenai/parser"
	"log"
)

// Run 爬虫项目的入口，输入种子数据，然后使用深度遍历算法进行爬取
func Run() {
	// 初始化种子数据
	var seeds []types.Request
	seeds = append(seeds, types.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// 遍历商品信息
	for _, seed := range seeds {
		// 调用fetcher，进行数据爬取
		body, err := fetcher.Fetch(seed.Url)
		if err != nil {
			log.Printf("")
		}

		// 调用parser，进行数据解析
		parseResult := seed.ParserFunc(body)
		seeds = append(seeds, parseResult...)

		// 返回结果进行入队
	}

}
