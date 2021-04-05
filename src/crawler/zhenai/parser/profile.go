// Package provides ...
package parser

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"hellosz.top/src/crawler/model"
	"hellosz.top/src/crawler/types"
)

const nickNameSelector = ".m-userInfo .info .nickName"
const idSelector = ".m-userInfo .info .id"
const infoSelector = ".m-userInfoFixed .CONTAINER .des"
const descSelector = ".CONTAINER .m-des span"

// ParseProfile 析用户信息
func ParseProfile(contents []byte) types.ParseResult {
	// 读取其中的节点
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		panic(err)
	}

	// 解析用户信息
	result := types.ParseResult{}
	profile := model.UserProfile{}

	// 昵称
	nickNameDom := extractFirstDom(dom, nickNameSelector)
	profile.Nickname = nickNameDom.Text()

	// ID
	idDom := extractFirstDom(dom, idSelector)
	profile.ID, err = strconv.Atoi(strings.TrimPrefix(idDom.Text(), "ID："))
	if err != nil {
		profile.ID = 0
	}
	fmt.Println(profile)

	// 基础信息
	infoDom := extractFirstDom(dom, infoSelector)
	info := infoDom.Text()
	infoArr := strings.Split(strings.ReplaceAll(info, " ", ""), "|")

	// for test
	profile.City = infoArr[0]
	profile.Age, err = strconv.Atoi(strings.TrimSuffix(infoArr[1], "岁"))
	if err != nil {
		log.Printf("wrong age desc: %s", infoArr[1])
		profile.Age = 0
	}
	profile.Education = infoArr[2]
	profile.Marriage = infoArr[3]
	profile.Height, err = strconv.Atoi(strings.TrimSuffix(infoArr[4], "cm"))
	if err != nil {
		log.Printf("wrong height desc: %s", infoArr[4])
		profile.Height = 0
	}
	profile.Salary = infoArr[5]

	// 描述
	descDom := extractFirstDom(dom, descSelector)
	profile.Description = descDom.Text()

	// 打印解析结果
	log.Printf("saving user profile:%s, %v\n", profile.Nickname, profile)

	return result
}

// 根据css选择器，返回解析的第一个节点
func extractFirstDom(document *goquery.Document, selector string) *goquery.Selection {
	return document.Find(selector).First()
}
