package persist

import (
	"encoding/json"
	"fmt"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/modules/log"
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
