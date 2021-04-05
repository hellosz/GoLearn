package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 读取本地缓存的数据结构
	html, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	// 调用ParseCityList获取返回结果
	parseResult := ParseCityList(html)
	requests := parseResult.Requests

	const resultCount int = 470
	requestUrls := []string{"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	requestItems := []string{"阿坝", "阿克苏", "阿拉善盟"}

	// test count
	if len(requests) != resultCount {
		t.Errorf("Expected got %d urls, but actually got %d", resultCount, len(requests))
	}

	if len(parseResult.Items) != resultCount {
		t.Errorf("Expected got %d items, but actually got %d", resultCount, len(parseResult.Items))
	}

	// test details
	for i, url := range requestUrls {
		if expectedUrl := parseResult.Requests[i].Url; url != expectedUrl {
			t.Errorf("Expected got url %s,but actually got %s", url, expectedUrl)
			t.Error(parseResult.Requests[i].Url)

		}
	}

	for i, item := range requestItems {
		if expectedItem := parseResult.Items[i]; item != expectedItem {
			t.Errorf("Expected got item %s,but actually got %s", item, expectedItem)
			t.Error(parseResult.Requests[i].Url)
		}
	}
}
