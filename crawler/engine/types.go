package engine

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

func NilParser(contents string) ParseResult {
	return ParseResult{}
}
