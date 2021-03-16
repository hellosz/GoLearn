package main

import "http/client"

func main() {
	// url := "http://localhost:8000/index"
	url := "https://www.imooc.com"
	client := client.GMCCrawler{}
	client.Crawler(url)
}
