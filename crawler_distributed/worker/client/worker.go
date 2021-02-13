package client

import (
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler_distributed/config"
	"douban-book-crawler/crawler_distributed/worker"
	"fmt"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			fmt.Println(err)
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
