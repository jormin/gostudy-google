package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	b, err := ioutil.ReadFile("zhenai.txt")
	if err != nil {
		panic(err)
	}
	parseResult := ParseCityList(string(b))
	const resutSize = 470
	expectUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectCitys := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	if len(parseResult.Requests) != resutSize {
		t.Errorf("Result should have %d requests; but had %d", resutSize, len(parseResult.Requests))
	}
	for i, url := range expectUrls {
		if parseResult.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, parseResult.Requests[i].Url)
		}
	}
	if len(parseResult.Items) != resutSize {
		t.Errorf("Result should have %d items; but had %d", resutSize, len(parseResult.Items))
	}
	for i, city := range expectCitys {
		if parseResult.Items[i].Name != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, parseResult.Items[i])
		}
	}
}
