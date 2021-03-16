package real

import "time"
import "net/http"
import "io/ioutil"

type HttpClient struct {
	UserAgent string
	TimeOut   time.Duration
}

func (h *HttpClient) Get(url string) string {
	// 获取请求
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 解析请求
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// 返回结果
	return string(body)
}
