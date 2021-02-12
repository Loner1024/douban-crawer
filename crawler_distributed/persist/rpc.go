package persist

import (
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler/persist"
	"log"

	"github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
	Client *elastic.Client
}

func (s ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v.", item, err)
	}
	return err
}
