package engine

import "fmt"

type Request struct {
	Url       string
	ParseFunc func(contents string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Tag  string
	Name string
	URL  string
	Data interface{}
}

func (i Item) String() string {
	return fmt.Sprintf("[TAG] %s [NAME] %s [URL] %s [DATA] %+v", i.Tag, i.Name, i.URL, i.Data)
}

func NilParser(contents string) ParseResult {
	return ParseResult{}
}
