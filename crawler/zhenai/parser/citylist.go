package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/modules/log"
	"strings"
)

// 解析城市列表
func ParseCityList(contents string) engine.ParseResult {
	result := engine.ParseResult{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(contents))
	if err != nil {
		log.Error("Parse city list error: %v", err)
		return result
	}
	dd := dom.Find(".city-list").Find("dd")
	dd.Each(func(i int, s *goquery.Selection) {
		s.Each(func(i int, sub *goquery.Selection) {
			sub.Find("a").Each(func(i int, a *goquery.Selection) {
				url, _ := a.Attr("href")
				result.Requests = append(result.Requests, engine.Request{
					Url:       url,
					ParseFunc: ParseUserList,
				})
				// 每个城市增加到10页
				for i := 2; i <= 10; i++ {
					result.Requests = append(result.Requests, engine.Request{
						Url:       fmt.Sprintf("%s/%d", url, i),
						ParseFunc: ParseUserList,
					})
				}
				result.Items = append(result.Items, engine.Item{
					Tag:  "city",
					Name: a.Text(),
					URL:  url,
				})
			})
		})
	})
	return result
}
