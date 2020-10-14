package worker

import (
	"github.com/jormin/go-study/crawler/engine"
)

type CrawlService struct {
}

func (c *CrawlService) Process(req Request, res *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineRes, err := engine.Parse(engineReq)
	if err != nil {
		return err
	}
	*res = SerializeResult(engineRes)
	return nil
}
