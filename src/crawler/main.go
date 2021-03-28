package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 获取城市列表页
	url := "https://www.zhenai.com/zhenghun"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	// 添加请求头(否则会40)
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// 打印结果
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error occured, status code: %d", res.StatusCode)
	}

	// 使用regexp读取数据

}
