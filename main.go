package main

import (
	"douban-book-crawler/douban/parser"
	"douban-book-crawler/engine"
	"douban-book-crawler/global"
	"douban-book-crawler/persist"
	"douban-book-crawler/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	concurrentEngine.Run(engine.Request{
		Url:        global.DoubanBookBaseUrl + "/tag",
		ParserFunc: parser.ParseTagList,
	})
}
