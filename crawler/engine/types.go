package engine

type Request struct {
	Url       string
	ParseFunc func(contents string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser(contents string) ParseResult {
	return ParseResult{}
}
