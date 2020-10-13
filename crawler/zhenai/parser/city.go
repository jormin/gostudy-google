package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/modules/log"
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
		//result.Requests = append(result.Requests, engine.Request{
		//	Url:       url,
		//	ParseFunc: ParseProfile,
		//})
		result.Items = append(result.Items, engine.Item{
			Tag:  "user",
			Name: nickname,
			URL:  url,
		})
	})

	dom.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, sub *goquery.Selection) {
			url, _ := sub.Attr("href")
			result.Requests = append(result.Requests, engine.Request{
				Url:       url,
				ParseFunc: ParseUserList,
			})
			result.Items = append(result.Items, engine.Item{
				Tag:  "Column",
				Name: sub.Text(),
				URL:  url,
			})
		})
	})
	return result
}
