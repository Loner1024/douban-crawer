package engine

import (
	"douban-book-crawler/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url, "")
	if err != nil {
		log.Printf("Worker error: %v\n", err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
