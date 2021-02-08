package parser

import (
	"douban-book-crawler/engine"
	"douban-book-crawler/global"
	"regexp"
)

const tagListRe string = `<a href="(/tag/[^>]*)">([^<]*)</a>`

func ParseTagList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(tagListRe)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, v := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url:        global.DoubanBookBaseUrl + string(v[1]),
			ParserFunc: ParseTag,
		})
		result.Item = append(result.Item, string(v[2]))
	}
	return result
}
