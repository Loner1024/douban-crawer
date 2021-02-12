package parser

import (
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler/model"
	"regexp"
	"strconv"
)

var authorRe = regexp.MustCompile(`<span>[\n|\s]+<span class="pl"> 作者</span>:[\n|\s]+<a class="" href=[^>]+>([^<]*)</a>[\n|\s]+</span>`)

var publishRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\s]([^<]*)<br/>`)
var publishYearRe = regexp.MustCompile(`<span class="pl">出版年:</span>[\s]([^<]*)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span>[\s]([^<]*)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>[\s]([^<]*)元<br/>`)
var isbnRe = regexp.MustCompile(`<span class="pl">ISBN:</span>[\s]([^<]*)<br/>`)
var summaryRe = regexp.MustCompile(`<div class="intro">[\n\s]+(([\s\S])*?)</div>`)
var idUrlRe = regexp.MustCompile(`https://book.douban.com/subject/([\d]+)/`)

// var authorSummaryRe = regexp.MustCompile()

func ParseBook(contents []byte, name string, url string) engine.ParseResult {
	var book = model.Book{
		Name:        name,
		Author:      extractString(contents, authorRe),
		Publisher:   extractString(contents, publishRe),
		PublishYear: extractString(contents, publishYearRe),
		Pages: func() int {
			page, _ := strconv.Atoi(extractString(contents, pageRe))
			return page
		}(),
		Price: func() float64 {
			price, err := strconv.ParseFloat(extractString(contents, priceRe), 64)
			if err != nil {
				return 0
			}
			return price
		}(),
		ISBN:          extractString(contents, isbnRe),
		Summary:       extractSummary(contents, summaryRe)[0],
		AuthorSummary: extractSummary(contents, summaryRe)[1],
	}
	return engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Id:      extractString([]byte(url), idUrlRe),
				Payload: book,
			},
		},
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	}
	return ""
}

func extractSummary(contents []byte, re *regexp.Regexp) []string {
	match := re.FindAllSubmatch(contents, -1)
	if len(match) > 1 {
		if len(match[0]) > 1 && len(match[1]) > 1 {
			return []string{string(match[0][1]), string(match[1][1])}
		}
	}
	return []string{"", ""}
}
