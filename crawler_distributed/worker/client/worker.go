package client

import (
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler_distributed/config"
	"douban-book-crawler/crawler_distributed/rpcsupport"
	"douban-book-crawler/crawler_distributed/worker"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			fmt.Println(err)
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
