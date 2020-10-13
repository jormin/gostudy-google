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
	"time"
)

var rateLimter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimter
	reqest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"cookie":     "FSSBBIl1UgzbN7NO=547Uts.YHCiT0M8IpsTSsYix6WsQDtppnekX.nznMQjUgunF5ozEusVH.jy6j28Mv2PN16ZCuPCsRCpSYnm0OQq; sid=21d3967b-a632-4775-8590-bae331ffbb96; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1602577638; ec=u5qdoTwo-1602577638200-0f359067110c6-951124028; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1602577827; _efmdata=0MmYZpLe0v%2FbvufnKWapQeFu0Xt7DqfK%2BfBhedKohzXHY4FANgNCDiZZMTWb4RNyFSNChMpS1%2Bhab96rF401R7%2FuYC9W%2BcYGDOqc%2BzvLrDs%3D; _exid=gqjDmsQWZt2%2FDyAPP8f7Fb%2B7XZreNMMuA%2BL3V333i8mZQoh91wKuVlNsBHnRCj2rb%2BVLo2HGSGL3LIjsn2d7qg%3D%3D; FSSBBIl1UgzbN7NP=5UgY2Nm5mRFEqqqmTEpwVnG.HKXDVCBmscFjQ46J9FTafN2oc7hJql98tPfpKkBlJuQLwlOGxnbkwEvbVw8T5SkZevVPJHDVAjUcYm5CevDo3ezoTK2h1BEz6kTPrcv9.kqlkAtrj7bsDPYxWld.xF2_giuZyc4hGPsyqdXbPot.OMOsX730n1bj2A.H3BlUt9XJW59jzMTrDL3btC_XDsXy0VzujOXz.yr77nJ6AH0evbgnB1UjXXXfiv3RZW2FS0",
		"referer":    "https://www.zhenai.com/",
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
	}
	for key, val := range headers {
		reqest.Header.Add(key, val)
	}
	client := http.Client{}
	resp, err := client.Do(reqest)
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
