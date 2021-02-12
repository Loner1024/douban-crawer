package model

import "encoding/json"

type Book struct {
	Name          string
	Author        string
	Publisher     string
	PublishYear   string
	Pages         int
	Price         float64
	ISBN          string
	Summary       string
	AuthorSummary string
}

func FromJsonObj(o interface{}) (Book, error) {
	var book Book
	s, err := json.Marshal(o)
	if err != nil {
		return book, err
	}
	err = json.Unmarshal(s, &book)
	return book, err
}
