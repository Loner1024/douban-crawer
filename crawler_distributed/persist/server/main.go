package main

import (
	"douban-book-crawler/crawler_distributed/persist"
	"douban-book-crawler/crawler_distributed/rpcsupport"
	"flag"
	"fmt"

	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	serveRpc(fmt.Sprintf(":%d", *port))
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
