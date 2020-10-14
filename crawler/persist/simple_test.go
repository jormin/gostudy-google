package persist

import (
	"context"
	"encoding/json"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/zhenai/parser"
	"github.com/jormin/go-study/helper"
	"github.com/jormin/go-study/modules/log"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"testing"
)

func TestSave(t *testing.T) {

	elasticClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		helper.LogError("Connect elasticsearch error", err)
		panic(err)
	}
	index := "profile"

	b, _ := ioutil.ReadFile("../zhenai/parser/city.html")
	result := parser.ParseUserList(string(b))
	for _, item := range result.Items {
		log.Info("%v", item)
		id, err := engine.Save(elasticClient, index, item)
		if err != nil {
			t.Errorf("Save error: %v", err)
			return
		}
		// todo try to start up elastic search
		// here use docker go client.
		client, err := elastic.NewClient(elastic.SetSniff(false))
		if err != nil {
			t.Errorf("Connect elasticsearch error: %v", err)
		}
		resp, err := client.Get().Index("profile").Id(id).Do(context.Background())
		if err != nil {
			t.Errorf("Index document error: %v", err)
		}
		actual, _ := resp.Source.MarshalJSON()
		expect, _ := json.Marshal(item)
		if string(actual) != string(expect) {
			t.Errorf("got %v; expected %v", actual, item)
		}
	}
}
