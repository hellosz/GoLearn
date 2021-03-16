package client

import (
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
)

type GMCCrawler struct {
	via bool
}

// func (g GMCCrawler)getVia() {
// }

func (g GMCCrawler) Crawler(url string) {
	// 简单的http客户端
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()
	// 复杂的客户端
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加请求头
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	// 发送请求
	// resp, err := http.DefaultClient.Do(request)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("redirect to :%v", req)

			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// // 使用httputil读取数据
	// content, err := httputil.DumpResponse(resp, false)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%q", content)

	// fmt.Println()

	// // 使用http读取数据
	// dump, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%q", dump)
	fmt.Print("OK")
}
