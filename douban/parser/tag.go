package parser

import (
	"douban-book-crawler/engine"
	"regexp"
)

var tagRe = regexp.MustCompile(`<a href="(https://book.douban.com/subject/[0-9]+/)"[^>]*>[\n|\s]*([^<]*)[\n|\s]*</a>`)

func ParseTag(contents []byte) engine.ParseResult {
	re := tagRe
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v := range match {
		name := string(v[2])
		url := string(v[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseBook(c, name, url)
			},
		})
	}
	return result
}
