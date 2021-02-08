package main

import (
	"douban-book-crawler/douban/parser"
	"douban-book-crawler/engine"
	"douban-book-crawler/global"
)

func main() {
	engine.Run(engine.Request{
		Url:        global.DoubanBookBaseUrl + "/tag",
		ParserFunc: parser.ParseTagList,
	})
}
