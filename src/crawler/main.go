package main

import (
	"fmt"
	"hellosz.top/src/crawler/engine"
	"hellosz.top/src/crawler/types"
	"hellosz.top/src/crawler/zhenai/parser"
	"io/ioutil"
)

func main() {
	// // 设置初始节点
	var seeds []types.Request
	seeds = append(seeds, types.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// 启动引擎，开始爬取
	engine.Run(seeds)
	// testProfile()
}

func testProfile() {
	profile, err := ioutil.ReadFile("./zhenai/parser/profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseProfile(profile)
	fmt.Println(result)
}
