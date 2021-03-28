package main

import client "hellosz.top/src/http/client"

func main() {
	// url := "https://www.imooc.com/"
	// url := "https://it.patpat.com/it/product/Adorable-Dinosaur-Driving-Print-Linen-Aprons-for-Mommy-and-Me-467904.html"
	// url := "https://us.patpat.com/product/Casual-Plaid-Splice-Hooded-Long-sleeve-Top-and-Pants-Set-for-Baby-378488.html"
	url := "https://it.patpat.com/it/product/Adorable-Dinosaur-Driving-Print-Linen-Aprons-for-Mommy-and-Me-467904.html?sku_id=19700466&currency=EUR&country_code=IT"
	client := client.GMCCrawler{}
	client.Crawler(url)
}
