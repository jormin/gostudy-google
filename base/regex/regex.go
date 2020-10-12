package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is xxx@gmail.com
email1 is xxx@qq.com
email2 is xxx@163.com
email3 is xxx@163.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
