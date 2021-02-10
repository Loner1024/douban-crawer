package fetcher

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header = map[string][]string{
		"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36"},
	}
	if err != nil {
		return nil, err
	}
	log.Printf("Fetching %v\n", url)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
