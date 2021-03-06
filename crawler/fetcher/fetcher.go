package fetcher

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// var rateLimiter = time.NewTicker(5 * time.Second)

func Fetch(Url, IP string) ([]byte, error) {
	_ = IP
	// proxyUrl, err := url.Parse(IP)
	// if err != nil {
	// 	return nil, err
	// }
	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyUrl),
	// }
	// client := &http.Client{
	// 	Transport: transport,
	// 	Timeout:   time.Second * 30,
	// }
	client := &http.Client{}
	req, err := http.NewRequest("GET", Url, nil)
	req.Header = map[string][]string{
		"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36"},
		"Cookie":     {generalRandomCookie()},
	}
	if err != nil {
		return nil, err
	}
	log.Printf("Fetching %v\n", Url)
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

func generalRandomCookie() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 11)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return "bid=" + string(b)
}
