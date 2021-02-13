package main

import (
	"douban-book-crawler/crawler/douban/parser"
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler/global"
	"douban-book-crawler/crawler/persist"
	"douban-book-crawler/crawler/scheduler"
)

func main() {
	// go func() {
	// 	for {
	// 		getProxyIP()
	// 		time.Sleep(time.Second / 5)
	// 	}
	// }()
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	concurrentEngine.Run(engine.Request{
		Url:    global.DoubanBookBaseUrl + "/tag/web",
		Parser: engine.NewFuncParser(parser.ParseTag, "parser.ParseTag"),
		// Parser: engine.NewFuncParser(parser.ParseTagList, "parser.ParseTagList"),
	})
}

// func getProxyIP() {
// 	resp, err := http.Get("")
// 	if err != nil {
// 		log.Printf("get proxy ip err %s", err)
// 	}
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("get proxy ip error, try restart")
// 			go getProxyIP()
// 		}
// 	}()
// 	defer resp.Body.Close()
// 	proxyReturn, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("read proxy ip err %s", err)
// 	}
// 	str := strings.Split(string(proxyReturn), "\n")
// 	for _, v := range str {
// 		global.ProxyIP <- "http://" + v
// 	}
// }
