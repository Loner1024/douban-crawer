package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

func NilParserFunc([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string
	Id      string
	Payload interface{}
}

type ParserFunc func(contents []byte, url string) ParseResult
