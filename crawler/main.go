package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	// 如果是GBK编码则需要转换
	utf8Reader := transform.NewReader(resp.Body, determineEncoding(resp.Body).NewDecoder())
	b, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	printCitiesAll(b)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	b, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}

func printCitiesAll(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`)
	match := re.FindAllSubmatch(contents, -1)
	for _, m := range match {
		fmt.Printf("City: %s, URL: %s", m[2], m[1])
		fmt.Println()
	}
	fmt.Printf("Matchs found: %d\n", len(match))
}
