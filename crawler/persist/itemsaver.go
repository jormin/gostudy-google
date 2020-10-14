package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/modules/log"
	"github.com/olivere/elastic/v7"
	"os"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		// write item to file
		file, err := os.OpenFile("user.txt", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
		defer file.Close()
		if err != nil {
			log.Error("open file error: %v", err)
		}
		count := 0
		for {
			item := <-out
			log.Info("Item saver got item #%d: %s", count, item)
			count++
			// Save
			b, _ := json.Marshal(item)
			c := fmt.Sprintf("%s\n", b)
			_, err := file.Write([]byte(c))
			if err != nil {
				log.Error("Save item error: %v", err)
			}
		}
	}()
	return out
}

func save(item engine.Item) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Error("Connect elasticsearch error: %v", err)
		return id, err
	}
	b, _ := json.Marshal(item)
	resp, err := client.Index().Index("dating_profile").BodyString(string(b)).Do(context.Background())
	if err != nil {
		log.Error("Index item error: %v", err)
		return id, err
	}
	return resp.Id, nil
}
