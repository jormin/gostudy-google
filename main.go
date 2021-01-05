package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v := get()
	fmt.Println(v)
}

func get() SkuActivity {
	list := "null"
	k := SkuActivity{}
	fmt.Println(k)
	fmt.Println(&k)
	var v SkuActivity
	_ = json.Unmarshal([]byte(list), &v)
	return v
}

type SkuActivity struct {
	SkuID            int64            `json:"sku_id"`
	SkuActivityCount int              `json:"sku_activity_count"`
	JoinActivityInfo JoinActivityInfo `json:"join_activity_info"`
}

type JoinActivityInfo struct {
	ActivityID    int64  `json:"activity_id"`
	Type          int    `json:"type"`
	ActivityName  string `json:"activity_name"`
	ActivityPrice int    `json:"activity_price"`
	ActivityStock int    `json:"activity_stock"`
	SkuID         int64  `json:"sku_id"`

	// todo 联系活动组更改返回数据参数类型
	ActivityStatus string `json:"activity_status"`
	StartAt        string `json:"start_at"`
	EndAt          string `json:"end_at"`
	OnceNum        string `json:"once_num"`
	BuyTimes       string `json:"buy_times"`

	CanPay    int `json:"can_pay"`
	SubStatus int `json:"sub_status"`
}
