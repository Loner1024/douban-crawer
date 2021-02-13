package main

import (
	"douban-book-crawler/crawler_distributed/config"
	"douban-book-crawler/crawler_distributed/rpcsupport"
	"douban-book-crawler/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "https://book.douban.com/subject/4913064/",
		Parser: worker.SerializedParser{
			Name: "BookParser",
			Args: "活着",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
