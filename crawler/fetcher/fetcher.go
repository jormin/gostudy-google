package fetcher

import (
	"bufio"
	"fmt"
	"github.com/jormin/go-study/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	// 如果是GBK编码则需要转换
	//bufioReader := bufio.NewReader(resp.Body)
	//e := determineEncoding(bufioReader)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(resp.Body)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	b, err := r.Peek(1024)
	if err != nil {
		log.Error("Fetch error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}
