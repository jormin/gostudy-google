package worker

import (
	"errors"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/zhenai/parser"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/helper"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Requests []Request
	Items    []engine.Item
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseUserList:
		return engine.NewFuncParser(parser.ParseUserList, config.ParseUserList), nil
	case config.ParseProfile:
		return engine.NewFuncParser(parser.ParseProfile, config.ParseProfile), nil
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unkonow parser name")
	}
}

func DeserializeRequest(r Request) (engine.Request, error) {
	deserializeParser, err := DeserializeParser(r.Parser)
	if err != nil {
		helper.LogError("Deserialize request error", err)
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: deserializeParser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}
