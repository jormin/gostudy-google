package persist

import (
	"encoding/json"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/model"
	"github.com/olivere/elastic/v7"
)

type SaverService struct {
	ElasticClient *elastic.Client
	Index         string
}

func (s *SaverService) Save(item engine.Item, res *string) error {
	b, _ := json.Marshal(item.Data)
	var profile model.SimpleProfile
	_ = json.Unmarshal(b, &profile)
	item.Data = profile
	id, err := engine.Save(s.ElasticClient, s.Index, item)
	if err == nil {
		*res = id
	}
	return err
}
