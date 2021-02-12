package main

import (
	"douban-book-crawler/crawler/douban/parser"
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler/global"
	"douban-book-crawler/crawler/scheduler"
	"douban-book-crawler/crawler_distributed/config"
	"douban-book-crawler/crawler_distributed/persist/client"
	"fmt"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan:    itemChan,
	}
	concurrentEngine.Run(engine.Request{
		Url:        global.DoubanBookBaseUrl + "/tag",
		ParserFunc: parser.ParseTagList,
	})
}
