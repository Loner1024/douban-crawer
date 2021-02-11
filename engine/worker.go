package engine

import (
	"douban-book-crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("worker error: %v\n", err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
