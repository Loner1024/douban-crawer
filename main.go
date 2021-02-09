package main

import (
	"douban-book-crawler/douban/parser"
	"douban-book-crawler/engine"
	"douban-book-crawler/global"
	"douban-book-crawler/scheduler"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	concurrentEngine.Run(engine.Request{
		Url:        global.DoubanBookBaseUrl + "/tag",
		ParserFunc: parser.ParseTagList,
	})
}
