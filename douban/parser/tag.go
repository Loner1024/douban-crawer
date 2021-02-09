package parser

import (
	"douban-book-crawler/engine"
	"regexp"
	"strings"
)

const tagRe string = `<a href="(https://book.douban.com/subject/[0-9]+/)"[^>]*>([^<]*)</a>`

func ParseTag(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(tagRe)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, v := range match {
		name := string(v[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(v[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseBook(c, name)
			},
		})
		result.Items = append(result.Items, strings.Replace(strings.Replace(name, "\n", "", -1), " ", "", -1))
	}
	return result
}
