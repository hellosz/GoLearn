package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const beforeScheme = "http://"
const afterScheme = "https://"

// 累计获取的url数量
var FetchedUrlCount = 0

// Fetch fetch data from url
func Fetch(url string) ([]byte, error) {
	// 输出日志当前获取的url数量
	log.Printf("FetchedUrlCount: %d", FetchedUrlCount)
	FetchedUrlCount++

	// 转换协议
	url = transSchema(url)
	log.Println(url)

	// 设置客户端
	log.Printf("is crawling url:%s\n", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// 可以自定义添加请求头
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")

	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	// 判断结果(针对album返回202，进行特殊处理) FIXME 解决登陆问题
	if resp.StatusCode == http.StatusAccepted || resp.StatusCode == http.StatusForbidden {
		log.Printf("%s return %d, try to read local file profile_test_data.html\n", url, resp.StatusCode)
		return ioutil.ReadFile("./zhenai/parser/profile_test_data.html")
	}

	// 判断结果
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong Status Code: %d", resp.StatusCode)
	}

	// 解析结果，并进行返回
	return ioutil.ReadAll(resp.Body)
}

// transSchema 转换协议，将http转换成https
func transSchema(url string) string {
	if strings.Contains(url, beforeScheme) {
		return strings.Replace(url, beforeScheme, afterScheme, 1)
	}

	return url
}
