# 豆瓣分布式爬虫

使用 Go 语言开发的豆瓣书籍分布式爬虫。以豆瓣图书 tag 页面为入口，抓取豆瓣图书信息。

分布式的各个 worker 之间使用 JSON-RPC 进行通信，抓取后将数据存储进入 ElasticSearch。

![未命名.001](./assets/img/未命名.001.jpeg)