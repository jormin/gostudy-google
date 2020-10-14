package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/model"
	"github.com/jormin/go-study/modules/log"
	"strconv"
	"strings"
)

// 解析用户列表
func ParseUserList(contents string) engine.ParseResult {
	result := engine.ParseResult{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(contents))
	if err != nil {
		log.Error("Parse city list error: %v", err)
		return result
	}
	dd := dom.Find(".g-list").Find(".list-item")
	dd.Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").Attr("href")
		nickname := s.Find(".content").Find("th").First().Text()
		profile := model.NewProfile()
		urlArr := strings.Split(url, "/")
		id, _ := strconv.Atoi(urlArr[len(urlArr)-1])
		profile.BasicInfo.ID = id
		profile.BasicInfo.Nickname = nickname
		tds := s.Find(".content").Find("td")
		tds.Each(func(i int, sub *goquery.Selection) {
			label := sub.Find("span").Text()
			val := strings.Replace(sub.Text(), label, "", 1)
			switch label {
			case "性别：":
				if val == "女士" {
					profile.NormalInfo.Gender = 1
				} else {
					profile.NormalInfo.Gender = 0
				}
			case "居住地：":
				profile.BasicInfo.City = val
			case "年龄：":
				profile.NormalInfo.Age = val
			case "学   历：":
				profile.NormalInfo.Education = val
			case "婚况：":
				profile.NormalInfo.Marital = val
			case "身   高：":
				profile.NormalInfo.Height = fmt.Sprintf("%scm", val)
			case "月   薪：":
				profile.NormalInfo.Salary = val
			}
		})
		profile.BasicInfo.Description = s.Find(".introduce").Text()
		//result.Requests = append(result.Requests, engine.Request{
		//	Url:       url,
		//	ParseFunc: ParseProfile,
		//})
		result.Items = append(result.Items, engine.Item{
			Tag:  "user",
			Name: nickname,
			URL:  url,
			Data: profile,
		})
	})

	// other city page
	dom.Find(".hot-city").Find(".city-list").Each(func(i int, cs *goquery.Selection) {
		if i == 0 || i == 3 {
			return
		}
		cs.Find(".list-item").Each(func(i int, s *goquery.Selection) {
			s.Find("a").Each(func(i int, sub *goquery.Selection) {
				url, _ := sub.Attr("href")
				result.Requests = append(result.Requests, engine.Request{
					Url:       url,
					ParseFunc: ParseUserList,
				})
				result.Items = append(result.Items, engine.Item{
					Tag:  "column",
					Name: sub.Text(),
					URL:  url,
				})
			})
		})
	})
	return result
}
