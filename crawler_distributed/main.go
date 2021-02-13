package main

import (
	"douban-book-crawler/crawler/douban/parser"
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler/global"
	"douban-book-crawler/crawler/scheduler"
	itemsaver "douban-book-crawler/crawler_distributed/persist/client"
	"douban-book-crawler/crawler_distributed/rpcsupport"
	worker "douban-book-crawler/crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

/*
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.11.0clear
*/

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workHosts     = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workHosts, ","))
	processor := worker.CreateProcessor(pool)
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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
