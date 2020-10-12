package fetcher

import (
	"bufio"
	"fmt"
	"github.com/jormin/go-study/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
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
	utf8Reader := transform.NewReader(resp.Body, determineEncoding(resp.Body).NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	b, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Error("Fetch error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}
