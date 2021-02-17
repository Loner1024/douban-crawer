package main

import (
	"context"
	"douban-book-crawler/crawler/engine"
	"reflect"

	"github.com/olivere/elastic/v7"
)

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{}
}

func QueryData(key string, start int) (SearchResult, error) {
	var result SearchResult
	result.Query = key

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return SearchResult{}, err
	}
	ctx := context.Background()
	termQuery := elastic.NewQueryStringQuery(key)
	searchResult, err := client.Search().Index("douban").Query(termQuery).From(start).Size(20).Do(ctx)
	if err != nil {
		return SearchResult{}, err
	}

	result.Hits = searchResult.TotalHits()
	result.Items = searchResult.Each(reflect.TypeOf(engine.Item{}))
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}
