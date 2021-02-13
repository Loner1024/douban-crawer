package main

import (
	"douban-book-crawler/crawler/douban/parser"
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler/global"
	"douban-book-crawler/crawler/scheduler"
	"douban-book-crawler/crawler_distributed/config"
	itemsaver "douban-book-crawler/crawler_distributed/persist/client"
	worker "douban-book-crawler/crawler_distributed/worker/client"
	"fmt"
)

/*
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.11.0clear
*/

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      50,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	concurrentEngine.Run(engine.Request{
		Url:    global.DoubanBookBaseUrl + "/tag",
		Parser: engine.NewFuncParser(parser.ParseTagList, "ParseTagList"),
	})
}
