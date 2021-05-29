package main

import (
	"fmt"

	client "hellosz.top/src/http/client"
)

func main() {
	// url := "https://www.imooc.com/"
	// url := "https://it.patpat.com/it/product/Adorable-Dinosaur-Driving-Print-Linen-Aprons-for-Mommy-and-Me-467904.html"
	// url := "https://us.patpat.com/product/Casual-Plaid-Splice-Hooded-Long-sleeve-Top-and-Pants-Set-for-Baby-378488.html"
	// url := "https://it.patpat.com/it/product/Adorable-Dinosaur-Driving-Print-Linen-Aprons-for-Mommy-and-Me-467904.html?sku_id=19700466&currency=EUR&country_code=IT"
	// url := "https://www.patpat.com/en/product/Casual-Colorblock-Long-sleeve-Nursing-Tee-443014.html?sku_id=19481960&currency=GBP&country_code=GB"
	// url := "https://uk.patpat.com/en/product/Casual-Colorblock-Long-sleeve-Nursing-Tee-443014.html?country_code=GB&currency=GBP&sku_id=19481960"
	// url := "https://uk.patpat.com/product/Baby-Toddler-Girl-Pretty-Floral-Print-Layered-Dresses-454228.html"
	// url := "https://uk.patpat.com/en/product/Two-piece-Family-Matching-Flower-Printed-Swimsuit-385423.html?sku_id=19013265&currency=GBP&country_code=GB"
	url := "https://fr.patpat.com/fr/product/Mosaic-Denim-Sitiching-Floral-Print-Family-Matching-Sets-473066.html?sku_id=19745779&currency=EUR&country_code=FR"
	client := client.GMCCrawler{}

	numbers := 20
	for i := 0; i < numbers; i++ {
		fmt.Printf("======================================请求第%d次\n", i+1)
		client.Crawler(url)
	}
}
