package worker

import (
	"douban-book-crawler/crawler/douban/parser"
	"douban-book-crawler/crawler/engine"
	"errors"
	"log"
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
	Items    []engine.Item
	Requests []Request
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

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineRequest, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("errpr deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineRequest)
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case "ParseTag":
		return engine.NewFuncParser(parser.ParseTag, "ParseTag"), nil
	case "ParseTagList":
		return engine.NewFuncParser(parser.ParseTagList, "ParseTagList"), nil
	case "BookParser":
		return parser.NewFuncBookParser(p.Args.(string)), nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
