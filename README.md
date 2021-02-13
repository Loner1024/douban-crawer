# 豆瓣分布式爬虫

```
├── crawler
│   ├── douban
│   │   └── parser
│   │       ├── book.go
│   │       ├── book_test.go
│   │       ├── book_test_data.html
│   │       ├── tag.go
│   │       ├── tag_test.go
│   │       ├── tag_test_data.html
│   │       ├── taglist.go
│   │       ├── taglist_test.go
│   │       └── taglist_test_data.html
│   ├── engine
│   │   ├── concurrent.go
│   │   ├── simple.go
│   │   ├── types.go
│   │   └── worker.go
│   ├── fetcher
│   │   └── fetcher.go
│   ├── global
│   │   └── global.go
│   ├── main.go
│   ├── model
│   │   └── douban_book.go
│   ├── persist
│   │   ├── itemsaver.go
│   │   └── itemsaver_test.go
│   ├── scheduler
│   │   ├── queued.go
│   │   └── simple.go
│   └── test
│       ├── main.go
│       └── test
├── crawler_distributed
│   ├── config
│   │   └── config.go
│   ├── main.go
│   ├── persist
│   │   ├── client
│   │   │   └── itemsaver.go
│   │   ├── rpc.go
│   │   └── server
│   │       ├── client_test.go
│   │       └── main.go
│   ├── rpcsupport
│   │   └── rpc.go
│   └── worker
│       ├── client
│       │   └── worker.go
│       ├── rpc.go
│       ├── server
│       │   ├── client_test.go
│       │   └── main.go
│       └── types.go
├── go.mod
└── go.sum
```