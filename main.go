package main

import (
	"douban-book-crawler/douban/parser"
	"douban-book-crawler/engine"
	"douban-book-crawler/global"
	"douban-book-crawler/persist"
	"douban-book-crawler/scheduler"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	concurrentEngine.Run(engine.Request{
		Url:        global.DoubanBookBaseUrl + "/tag",
		ParserFunc: parser.ParseTagList,
	})
}
