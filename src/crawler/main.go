package main

import (
	"fmt"
	"hellosz.top/src/crawler/engine"
	"hellosz.top/src/crawler/schedule"
	"hellosz.top/src/crawler/types"
	"hellosz.top/src/crawler/zhenai/parser"
	"io/ioutil"
)

func main() {
	// 设置初始节点
	var seeds []types.Request
	seeds = append(seeds, types.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// 多进程引擎
	engine := engine.Concurrent{
		Scheduler:   &schedule.Queue{},
		WorkerCount: 100,
	}
	engine.Run(seeds)

	// 单进程引擎
	// engine := engine.Simple{}
	// engine.Run(seeds)
}

func testProfile() {
	profile, err := ioutil.ReadFile("./zhenai/parser/profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseProfile(profile)
	fmt.Println(result)
}
