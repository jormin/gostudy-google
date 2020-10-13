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
		result.Requests = append(result.Requests, engine.Request{
			Url:       url,
			ParseFunc: ParseProfile,
		})
		result.Items = append(result.Items, nickname)
	})
	return result
}
