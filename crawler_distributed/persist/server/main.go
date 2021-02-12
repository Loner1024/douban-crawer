package main

import (
	"douban-book-crawler/crawler_distributed/config"
	"douban-book-crawler/crawler_distributed/persist"
	"douban-book-crawler/crawler_distributed/rpcsupport"
	"fmt"

	"github.com/olivere/elastic/v7"
)

func main() {
	serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort))
}

func serveRpc(host string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
	})
}
