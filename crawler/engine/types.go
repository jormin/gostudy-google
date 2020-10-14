package engine

type Parser interface {
	Parse(contents string) ParseResult
	Serialize() (name string, args interface{})
}

type ParseFunc func(contents string) ParseResult

type Request struct {
	Url    string
	Parser Parser
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

type NikParser struct {
}

func (p NikParser) Parse(contents string) ParseResult {
	return ParseResult{}
}

func (p NikParser) Serialize() (name string, args interface{}) {
	return "NikParser", nil
}

type FuncParser struct {
	ParseFun ParseFunc
	Name     string
}

func (p *FuncParser) Parse(contents string) ParseResult {
	return p.ParseFun(contents)
}

func (p *FuncParser) Serialize() (name string, args interface{}) {
	return p.Name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		ParseFun: p,
		Name:     name,
	}
}
