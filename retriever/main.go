package main

import (
	"fmt"
	"github.com/jormin/go-study/retriever/mock"
	"github.com/jormin/go-study/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPost interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(p Poster) {
	p.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func session(rp RetrieverPost) string {
	rp.Post(url, map[string]string{
		"contents": "another faked immoc.com",
	})
	return rp.Get(url)
}

func main() {
	var r Retriever
	retriever := &mock.Retriever{
		Contents: "this is a fake imooc.com",
	}
	insepect(retriever)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	insepect(r)

	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.Timeout)
	} else {
		fmt.Println("not a real retriever")
	}
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println(session(retriever))

	//fmt.Println(download(realRetriever))
	//fmt.Println(download(realRetriever))
}

func insepect(r Retriever) {
	fmt.Println("Inspecting")
	fmt.Printf("> %T,%v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
