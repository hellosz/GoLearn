package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch fetch data from url
func Fetch(url string) ([]byte, error) {
	// 设置客户端
	log.Printf("is crawling url:%s\n", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// 可以自定义添加请求头
	// request.Header.Add("User-Agent", "")
	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	// 判断结果
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong Status Code: %d", resp.StatusCode)
	}

	// 解析结果，并进行返回
	return ioutil.ReadAll(resp.Body)
}
