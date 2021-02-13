package main

import (
	"douban-book-crawler/crawler_distributed/config"
	"douban-book-crawler/crawler_distributed/rpcsupport"
	"douban-book-crawler/crawler_distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(
		rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}),
	)
}
