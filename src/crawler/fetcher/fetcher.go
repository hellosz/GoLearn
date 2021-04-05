package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const beforeScheme = "http://"
const afterScheme = "https://"

// Fetch fetch data from url
func Fetch(url string) ([]byte, error) {
	// 转换协议
	url = transSchema(url)
	fmt.Println(url)

	// 设置客户端
	log.Printf("is crawling url:%s\n", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// 可以自定义添加请求头
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")
	request.Header.Add("Cookie", "enable_FSSBBIl1UgzbN7N=true; Hm_ck_1617327893671=is-cookie-enabled; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1617327894; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1617150142; _efmdata=RsaBq7PEtQ6b8W%2FyHTS6bpQD0MiMAnX6kZMwE9URSxro9a8Uq1%2FF%2FjzpKqkvSc%2FeUt3KK6x29JPg5o8SOj4ZZm9can%2F0dxIAdHLsVHkeQlQ%3D; _exid=t6GENUjgG0O7qi%2Fiai9dDkLsJhEuauxM4TMn3HCh4g%2FNvIlUcFGM33i4cCWC6MIgu1eQIDQg3r7nZgyBje8fkQ%3D%3D; _pc_myzhenai_memberid_=%22%2C1244676710%22; _pc_myzhenai_showdialog_=1; ec=ZXRaarpJ-1617150132180-ed7948d2e36ea-1860306439; recommendId=%7B%22main-flow%22%3A%22baseline%22%2C%22off-feat%22%3A%22v1%22%2C%22feat-config%22%3A%22v1%22%2C%22model-version%22%3A%22v11%22%7D; refreshToken=1244676710.1617413826443.fca6557038e1e57ba53c9932385a0481; sid=4b46588e-f364-4667-a8cd-1af3c2849de0; spliterabparams=1617327345603%3A4772546598278732121; token=1244676710.1617327426442.e894393e2bee4e9eb8af289ee4ee08e9; FSSBBIl1UgzbN7NO=5HXfaqh.l6gD76vq9Sh4LURUiheDK8tXpgRD9n5KDM9bcfCttVQe1uwp0NilTOhJcAhKTiFjzrWtiyb8sggFdBG; FSSBBIl1UgzbN7NP=53DR6DDrbwSVqqqmg9lpmYGn8Qxa5UhBiQAm6rgdojnlbWwTeYbqm0gS.Q_oiTtwWg55kAYzyweO29vkaMOaB2GEf8dQypC.MdYMi4csPNSHTq_gUwW0q2TZdsaK1h6e.sG74AWwL58oKaiwkVI9j.W.bs4edZXETbS7Ykz3y52dgrtZgKldy7Qm88YrFeRKWeagR2DUoKMIjVsJGwXdKhqJm6fxrd.KRm.oTNh9Sb5tItd7ggq_k96BSmZHScXNyW")

	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	// 判断结果
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong Status Code: %d", resp.StatusCode)
	}

	// 进行编码转换
	bodyReader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(bodyReader, simplifiedchinese.GBK.NewDecoder())
	fmt.Println(ioutil.ReadAll(utf8Reader))

	// 解析结果，并进行返回
	return ioutil.ReadAll(utf8Reader)
}

// transSchema 转换协议，将http转换成https
func transSchema(url string) string {
	if strings.Contains(url, beforeScheme) {
		return strings.Replace(url, beforeScheme, afterScheme, 1)
	}

	return url
}
