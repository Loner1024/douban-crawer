package parser

import (
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler/global"
	"regexp"
)

const tagListRe string = `<a href="(/tag/[^>]*)">([^<]*)</a>`

func ParseTagList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(tagListRe)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url:    global.DoubanBookBaseUrl + string(v[1]),
			Parser: engine.NewFuncParser(ParseTag, "ParseTag"),
		})
		// result.Requests = append(result.Items, string(v[2]))
	}
	return result
}
