package client

import (
	"fmt"
	"io/ioutil"
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
	// request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	// request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.87 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	// request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")

	// request.Header.Add("User-Agent", "PostmanRuntime/7.26.8")
	// fmt.Println("请求网站")
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	fmt.Println("请求h5")

	// 发送请求
	// resp, err := http.DefaultClient.Do(request)
	redirectNumber := 0
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("redirect to :%v\n", req)

			redirectNumber++

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

	// 使用http读取数据
	dump, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", dump)
	fmt.Print("OK\n")
	fmt.Printf("跳转%d次\n", redirectNumber)
}
